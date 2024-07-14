package querys

import (
	"context"
	"log"

	"social-network/internal/models"
	"social-network/internal/views/websocket/types"
)

func AddToNotificationTable(userID string, notificationType string, relatedID int) error {
	query := `
	INSERT INTO 
			notifications (user_id, type, related_id, status) 
	VALUES 
			($1, $2, $3, $4)`
	_, err := DB.Exec(context.Background(), query, userID, notificationType, relatedID, "pending")
	if err != nil {
		log.Printf("database: Failed to add notification: %v", err)
		return err // Return error if failed to insert post
	}
	return nil
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
	request := models.Request{ID: requestID}
	query := `SELECT sender_id, receiver_id, status FROM follow_requests WHERE request_id = $1`
	err := DB.QueryRow(context.Background(), query, requestID).Scan(&request.Sender, &request.Receiver, &request.Status)
	if err != nil {
		log.Println("Failed to get follow request")
		return nil, err
	}
	return &request, nil
}

func GetUserNotifications(userID string) ([]types.Notification, error) {
	var notifications []types.Notification
	query := `
	SELECT 
			type,
			related_id
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
	for rows.Next() {
		var notificationType string
		var relatedID int
		rows.Scan(&notificationType, &relatedID)
		switch notificationType {
		case "follow_request":
			request, err := GetFollowRequest(relatedID)
			if err != nil {
				log.Println("Failed to get follow request")
				continue
			}
			followNotification, err := GetFollowRequestNotification(*request)
			if err != nil {
				log.Println("Failed to get follow request")
				continue
			}
			notifications = append(notifications, *followNotification)
		case "group_invite":
			invitationNotification, err := GetGroupInvitationData(userID, relatedID)
			if err != nil {
				log.Println("Failed to get invitation Data")
				continue
			}
			notifications = append(notifications, *invitationNotification)
		case "join_request":
			groupRequestNotification, err := GetGroupRequestData(userID, relatedID)
			if err != nil {
				log.Println("Failed to get group request Data")
				continue
			}
			notifications = append(notifications, *groupRequestNotification)
		case "event_notification":
			eventNotification, err := GetGroupEventData(userID, relatedID)
			if err != nil {
				log.Println("Failed to get group event Data")
				continue
			}
			notifications = append(notifications, *eventNotification)
		}
	}
	return notifications, err
}

func OrganizeFollowRequest(recieverUsername string, sender models.UserProfile) types.Notification {
	notification := types.Notification{
		Type:    "FOLLOW_REQUEST",
		Message: "You have a new follow request",
		ToUser:  recieverUsername,
		Metadata: types.FollowRequestMetadata{
			UserDetails: types.UserDetails{
				Username:  sender.Username,
				FirstName: sender.FirstName,
				LastName:  sender.LastName,
			},
		},
	}
	return notification
}

func OrganizeGroupRequest(groupCreator string, GroupTitle string, groupID int, requester models.UserProfile) types.Notification {
	notification := types.Notification{
		Type:    "REQUEST_TO_JOIN_GROUP",
		Message: "You have a new group request",
		ToUser:  groupCreator,
		Metadata: types.GroupRequestMetadata{
			UserDetails: types.UserDetails{
				Username:  requester.Username,
				FirstName: requester.FirstName,
				LastName:  requester.LastName,
			},
			Group: types.GroupNotification{
				ID:    groupID,
				Title: GroupTitle,
			},
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

func OrganizeGroupInvitation(recieverUsername string, groupID int, groupTitle string)types.Notification{
	notification := types.Notification{
		Type: "GROUP_INVITATION",
		Message: "You have a new group invitation",
		ToUser: recieverUsername,
					Metadata: types.GroupNotification{
									ID: groupID,
									Title: groupTitle,
					},
}
return notification
}