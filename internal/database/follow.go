package database

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

func CreateFollowRequest(request models.Request)(int, error){
	var requestID int
	query := `
	SELECT 
		 request_id
		FROM
			follow_requests
		WHERE
		 (sender_id = $1 AND receiver_id = $2) 
`
err := DB.QueryRow(context.Background(),query, request.SenderID, request.ReceiverID).Scan(requestID)
if err != nil && err.Error() != "no rows in result set" {
	log.Printf("database: Failed check for request: %v", err)
	return 0, err // Return error if failed to insert post
}
if request.Status == "pending"{
	log.Printf("request already made")
	return request.ID, nil 
}
	query = `
	INSERT INTO 
			follow_requests (sender_id, receiver_id) 
	VALUES 
			($1, $2)
	RETURNING request_id`
err = DB.QueryRow(context.Background(),query, request.SenderID, request.ReceiverID).Scan(&request.ID)
if err != nil {
	log.Printf("database: Failed to insert request into database: %v", err)
	return 0, err // Return error if failed to insert post
}
return request.ID, nil
}

func RespondToFollow(response models.Response) error {
	log.Print(response)
	if response.Status == "accepted" {
		query := `UPDATE follow_requests SET status = $1 WHERE sender_id = $2 AND receiver_id = $3 RETURNING request_id`
		err := DB.QueryRow(context.Background(),query,"accepted", response.FollowerID, response.FolloweeID).Scan(&response.ID)
		if err != nil {
			log.Printf("database: Failed to update response in database: %v", err)
			return err // Return error if failed to insert post
		} 
		if response.ID == 0 {
			log.Print("no match")
			return nil
		}
		query = `INSERT INTO 
			follower (follower_id, followed_id) 
	VALUES 
			($1, $2)`
	_, err = DB.Exec(context.Background(),query, response.FollowerID, response.FolloweeID)
			if err != nil {
				log.Printf("database: Failed to add follower: %v", err)
				return err // Return error if failed to insert post
			} 
	} else if response.Status == "rejected"{
		query := `UPDATE follow_requests SET status = 'rejected' WHERE sender_id = $1 AND receiver_id = $2`
		_, err := DB.Exec(context.Background(),query, response.FollowerID, response.FolloweeID)
		if err != nil {
			log.Printf("database: Failed to update response in database: %v", err)
			return err // Return error if failed to insert post
		} 
	}
	return nil
}