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
