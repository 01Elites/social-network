package querys

import (
	"context"
	"errors"
	"log"
	"time"

	"social-network/internal/models"
	"social-network/internal/views/websocket/types"
)

func AddToNotificationTable(userID string, notificationType string, relatedID int) (int, error) {
	var notificationID int
	query := `
	INSERT INTO 
			notifications (user_id, type, related_id, status) 
	VALUES 
			($1, $2, $3, $4)
	RETURNING notification_id`
	err := DB.QueryRow(context.Background(), query, userID, notificationType, relatedID, "pending").Scan(&notificationID)
	if err != nil {
		log.Printf("database: Failed to add notification: %v", err)
		return 0, err // Return error if failed to insert post
	}
	return notificationID, nil
}

func CancelNotification(relatedID int, notificationType string, userID string) error {
	query := `UPDATE notifications SET status = 'canceled' WHERE (related_id = $1 AND type = $2 AND user_id = $3)`
	_, err := DB.Exec(context.Background(), query, relatedID, notificationType, userID)
	if err != nil {
		log.Printf("database: Failed to update response in database: %v", err)
		return err // Return error if failed to insert post
	}
	return nil
}

func UpdateNotificationTable(relatedID int, status string, notificationType string, userID string) error {
	query := `UPDATE notifications SET status = $1, read = true WHERE (related_id = $2 AND type = $3 AND user_id = $4 AND status = $5)`
	_, err := DB.Exec(context.Background(), query, status, relatedID, notificationType, userID, "pending")
	if err != nil {
		log.Printf("database: Failed to update response in database: %v", err)
		return err // Return error if failed to insert post
	}
	return nil
}

func GetFollowRequest(requestID int) (*models.Request, error) {
	var count int
	request := models.Request{ID: requestID}
	query := `SELECT sender_id, receiver_id, created_at FROM follow_requests WHERE request_id = $1 AND status = 'pending'`
	err := DB.QueryRow(context.Background(), query, requestID).Scan(&request.Sender, &request.Receiver, &request.CreatedAt)
	if err != nil {
		log.Println("Failed to get follow request, test")
		return nil, err
	}
	query = `SELECT COUNT(*) FROM follower WHERE follower_id = (SELECT user_id FROM "user" WHERE user_name=$1)
	 AND followed_id = (SELECT user_id FROM "user" WHERE user_name=$2)`
	err = DB.QueryRow(context.Background(), query, request.Sender, request.Receiver).Scan(&count)
	if err != nil {
		log.Print(err)
		log.Println("Failed to get follow request")
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("user already following")
	}
	return &request, nil
}

func GetUserNotifications(userID string) ([]types.Notification, error) {
	var notifications []types.Notification
	query := `
	SELECT 
			notification_id,
			type,
			related_id,
			read
	FROM
			notifications
	WHERE
			user_id = $1
			AND status = 'pending';
`
	rows, err := DB.Query(context.Background(), query, userID)
	if err != nil && err.Error() != "no rows in result set" {
		log.Printf("database: Failed check for request: %v", err)
		return nil, err
	}
	var count int
	for rows.Next() {
		var notificationID int
		var notificationType string
		var relatedID int
		var read bool
		count++
		err := rows.Scan(&notificationID, &notificationType, &relatedID, &read)
		if err != nil {
			log.Printf("database: Failed to scan notification row: %v", err)
			continue
		}
		var notification *types.Notification
		switch notificationType {
		case "follow_request":
			request, err := GetFollowRequest(relatedID)
			if err != nil && err.Error() != "no rows in result set" {
				log.Println("Failed to get follow request")
				continue
			}
			notification, err = GetFollowRequestNotification(*request)
			if err != nil && err.Error() != "no rows in result set" {
				log.Println("Failed to get follow request")
				continue
			}
		case "group_invite":
			notification, err = GetGroupInvitationData(userID, relatedID)
			if err != nil && err.Error() != "no rows in result set" {
				log.Println("Failed to get invitation Data")
				continue
			}

		case "join_request":
			notification, err = GetGroupRequestData(userID, relatedID)
			if err != nil && err.Error() != "no rows in result set" {
				log.Println("Failed to get group request Data")
				continue
			}

		case "event_notification":
			notification, err = GetGroupEventData(userID, relatedID)
			if err != nil && err.Error() != "no rows in result set" {
				log.Println("Failed to get group event Data")
				continue
			}
			log.Println("Follow request notification")

		default:
			log.Printf("Unknown notification type: %s", notificationType)
			continue
		}
		if notification != nil {
			notification.ID = notificationID
			notification.Read = read
			notifications = append(notifications, *notification)
		}
	}
	if err = rows.Err(); err != nil {
		log.Printf("database: Failed during row iteration: %v", err)
		return nil, err
	}
	log.Print(count)
	return notifications, err
}

func OrganizeFollowRequest(recieverUsername string, sender models.UserProfile, createdAt time.Time) types.Notification {
	notification := types.Notification{
		Type:    "FOLLOW_REQUEST",
		Message: "You have a new follow request",
		ToUser:  recieverUsername,
		Metadata: types.FollowRequestMetadata{
			UserDetails: types.UserDetails{
				Username:  sender.Username,
				FirstName: sender.FirstName,
				LastName:  sender.LastName,
				Avatar:    sender.Avatar,
			},
			CreationDate: createdAt,
		},
	}
	return notification
}

func OrganizeGroupRequest(groupCreator string, GroupTitle string, groupID int, requester models.UserProfile, createdAt string) types.Notification {
	notification := types.Notification{
		Type:    "REQUEST_TO_JOIN_GROUP",
		Message: "You have a new group request",
		ToUser:  groupCreator,
		Metadata: types.GroupRequestNotification{
			Requester: models.Requester{
				User: models.PostFeedProfile{
					UserName:  requester.Username,
					FirstName: requester.FirstName,
					LastName:  requester.LastName,
					Avatar:    requester.Avatar,
				},
				CreationDate: createdAt,
			},
			ID:    groupID,
			Title: GroupTitle,
		},
	}
	return notification
}

func OrganizeGroupEventRequest(member string, groupTitle string, groupID int, groupEvent types.EventDetails) types.Notification {
	notification := types.Notification{
		Type:    "EVENT",
		Message: "You have a new event in the group",
		ToUser:  member,
		Metadata: types.GroupEventMetadata{
			Group: types.GroupNotification{
				ID:    groupID,
				Title: groupTitle,
			},
			Event: groupEvent,
		},
	}
	return notification
}

func OrganizeGroupInvitation(recieverUsername string, groupID int, groupTitle string, invitedBy models.Requester) types.Notification {
	notification := types.Notification{
		Type:    "GROUP_INVITATION",
		Message: "You have a new group invitation",
		ToUser:  recieverUsername,
		Metadata: types.GroupNotification{
			ID:        groupID,
			Title:     groupTitle,
			InvitedBy: invitedBy,
		},
	}
	return notification
}

// SetNotificationAsRead sets the notification with the given ID as read in the database
func SetNotificationAsRead(notificationID int) error {
	query := `UPDATE notifications SET read = TRUE WHERE notification_id = $1`
	_, err := DB.Exec(context.Background(), query, notificationID)
	if err != nil {
		log.Printf("database: Failed to update notification in database: %v", err)
		return err // Return error if failed to update notification
	}
	return nil
}
