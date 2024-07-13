package querys

import (
	"context"
	"log"
	"social-network/internal/models"
	"social-network/internal/views/websocket/types"
)

func GetUsersFollowingByID(userID string) (map[string]bool, error) {
	Following := make(map[string]bool)
	query := `SELECT followed_id FROM follower WHERE follower_id = $1`
	rows, err := DB.Query(context.Background(), query, userID)
	if err != nil {
		log.Printf("database failed to scan followed user: %v\n", err)
		return nil, err
	}
	for rows.Next() {
		var followedUserID string
		if err = rows.Scan(&followedUserID); err != nil {
			log.Printf("database failed to scan followed user: %v\n", err)
			return nil, err
		}
		Following[followedUserID] = true
	}
	return Following, nil
}

func GetUserFollowingUserNames(userID string) ([]string, error) {
	var following []string
	query := `
	SELECT
    "user".user_name
	FROM
			follower
	INNER JOIN
			public."user" ON follower.followed_id = public."user".user_id
	WHERE
			follower.follower_id = $1
	`
	rows, err := DB.Query(context.Background(), query, userID)
	if err != nil {
		log.Printf("database failed to scan followed user: %v\n", err)
		return nil, err
	}
	for rows.Next() {
		var followedUsername string
		if err = rows.Scan(&followedUsername); err != nil {
			log.Printf("database failed to scan followed user: %v\n", err)
			return nil, err
		}
		following = append(following, followedUsername)
	}
	return following, nil

}


func GetUserFollowerUserNames(userID string) ([]string, error) {
	var followers []string
	query := `
	SELECT
		"user".user_name
	FROM
		follower
	INNER JOIN
		public."user" ON follower.follower_id = public."user".user_id
	WHERE
		follower.followed_id = $1
	`
	rows, err := DB.Query(context.Background(), query, userID)
	if err != nil {
		log.Printf("database failed to scan followers: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var followerUsername string
		if err = rows.Scan(&followerUsername); err != nil {
			log.Printf("database failed to scan follower user: %v\n", err)
			return nil, err
		}
		followers = append(followers, followerUsername)
	}

	if err = rows.Err(); err != nil {
		log.Printf("rows iteration error: %v\n", err)
		return nil, err
	}

	return followers, nil
}


func GetUsersFollowees(userID string) (map[string]bool, error) {
	Followees := make(map[string]bool)
	query := `SELECT follower_id FROM follower WHERE followed_id = $1`
	rows, err := DB.Query(context.Background(), query, userID)
	if err != nil {
		log.Printf("database failed to scan follower user: %v\n", err)
		return nil, err
	}
	for rows.Next() {
		var followerUserID string
		if err = rows.Scan(&followerUserID); err != nil {
			log.Printf("database failed to scan follower user: %v\n", err)
			return nil, err
		}
		Followees[followerUserID] = true
	}
	return Followees, nil
}

func CreateFollowRequest(request models.Request) (int, error) {
	// Check if request already exists
	query := `
		SELECT 
				request_id, status
		FROM
				follow_requests
		WHERE
				sender_id = $1 
				AND receiver_id = $2
				AND status NOT IN ('accepted', 'rejected' , 'canceled');
`
	err := DB.QueryRow(context.Background(), query, request.Sender, request.Receiver).Scan(&request.ID, &request.Status)
	if err != nil && err.Error() != "no rows in result set" {
		log.Printf("database: Failed check for request: %v", err)
		return 0, err // Return error if failed to insert post
	}

	// if it is already sent cancel the request by updating the status to canceled
	if request.Status == "pending" {
		query := `UPDATE follow_requests SET status = 'canceled' WHERE sender_id = $1 AND receiver_id = $2 AND request_id = $3`
		err := DB.QueryRow(context.Background(), query, request.Sender, request.Receiver, request.ID).Scan(&request.ID)
		if err != nil && err.Error() != "no rows in result set" {
			log.Printf("database1: Failed to update request in database: %v", err)
			return 0, err
		}
		return 0, nil
	}

	// Insert request into database
	query = `
	INSERT INTO 
			follow_requests (sender_id, receiver_id) 
	VALUES 
			($1, $2)
	RETURNING request_id`
	err = DB.QueryRow(context.Background(), query, request.Sender, request.Receiver).Scan(&request.ID)
	if err != nil {
		log.Printf("database: Failed to insert request into database: %v", err)
		return 0, err // Return error if failed to insert post
	}
	return request.ID, nil
}

func RespondToFollow(response models.Response) error {
	query := `UPDATE follow_requests SET status = $1 WHERE sender_id = $2 AND receiver_id = $3 RETURNING request_id`
	err := DB.QueryRow(context.Background(), query, response.Status, response.Follower, response.Followee).Scan(&response.ID)
	if err != nil {
		log.Printf("database: Failed to update response in database: %v", err)
		return err // Return error if failed to insert post
	}
	if response.ID == 0 {
		log.Print("no match")
		return nil
	}
	if response.Status == "accepted" {
		query = `INSERT INTO 
			follower (follower_id, followed_id) 
	VALUES 
			($1, $2)`
		_, err = DB.Exec(context.Background(), query, response.Follower, response.Followee)
		if err != nil {
			log.Printf("database: Failed to add follower: %v", err)
			return err // Return error if failed to insert post
		}
	}
	return nil
}

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

func UpdateNotificationTable(relatedID int, status string, notificationType string, userID string) error {
	query := `UPDATE notifications SET status = $1, read = true WHERE (related_id = $2 AND type = $3 AND user_id = $4 AND status = $5)`
	_, err := DB.Exec(context.Background(), query, status, relatedID, notificationType, userID, "pending")
	if err != nil {
		log.Printf("database: Failed to update response in database: %v", err)
		return err // Return error if failed to insert post
	}
	return nil
}

func GetUserNotifications(userID string)([]types.Notification,error) {
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
		return nil,err
		}
	for rows.Next(){
		var notificationType string
		var relatedID int
		rows.Scan(&notificationType, &relatedID)
		switch notificationType{
		case "follow_request":
			
		case "group_invite":

		case "join_request":
			notifications = append(notifications, OrganizeGroupRequest(GetGroupRequestData(userID, relatedID)))
		case "event_notification":
			notifications = append(notifications,OrganizeGroupEventRequest(GetGroupEventData(userID, relatedID)))
		}
	}
	return notifications, err
}

func OrganizeGroupRequest(groupCreator string, GroupTitle string, groupID int, requester models.UserProfile)types.Notification{
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

func OrganizeGroupEventRequest(member string, groupTitle string, groupID int, groupEvent types.EventDetails)types.Notification {
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