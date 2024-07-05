package querys

import (
	"context"
	"log"
	"social-network/internal/models"
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

func CreateFollowRequest(request models.Request) (int, error) {
	var requestID int
	query := `
	SELECT 
		 request_id
		FROM
			follow_requests
		WHERE
		 (sender_id = $1 AND receiver_id = $2) 
`
	err := DB.QueryRow(context.Background(), query, request.Sender, request.Receiver).Scan(requestID)
	if err != nil && err.Error() != "no rows in result set" {
		log.Printf("database: Failed check for request: %v", err)
		return 0, err // Return error if failed to insert post
	}
	if request.Status == "pending" {
		log.Printf("request already made")
		return request.ID, nil
	}
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
	err := DB.QueryRow(context.Background(), query, response.Status, response.FollowerID, response.FolloweeID).Scan(&response.ID)
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
		_, err = DB.Exec(context.Background(), query, response.FollowerID, response.FolloweeID)
		if err != nil {
			log.Printf("database: Failed to add follower: %v", err)
			return err // Return error if failed to insert post
		}
	}
	return nil
}

// IsFollowing checks if a user is following another user
func IsFollowing(userID string, followedID string) bool {
	query := `SELECT followed_id FROM follower WHERE follower_id = $1 AND followed_id = $2`
	err := DB.QueryRow(context.Background(), query, userID, followedID).Scan(&followedID)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return false
		}
		log.Printf("database: Failed to check if user is following: %v", err)
		return false
	}
	return true
}

func AddToNotificationTable() {}

func UpdateNotificationTable(notificationID int, status string, userID string) error {
	query := `UPDATE notifications SET status = $1 AND SET read = TRUE WHERE notification_id = $2`
	_, err := DB.Exec(context.Background(), query, status, notificationID)
	if err != nil {
		log.Printf("database: Failed to update response in database: %v", err)
		return err // Return error if failed to insert post
	}

	return nil
}

func ViewNotificationTable() {}
