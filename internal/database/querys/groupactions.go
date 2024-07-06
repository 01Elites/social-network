package querys

import (
	"context"
	"log"
	"social-network/internal/models"
	// "github.com/jackc/pgx/v5"
)

func CreateInvite(groupID int, senderID string, receiverID string) (int, error) {
	var status string
	var invitationID int
	query := `
	SELECT 
		 status,
		 invitation_id
		FROM
			group_invitations
		WHERE
		 (group_id = $1 AND receiver_id = $2) 
`
	err := DB.QueryRow(context.Background(), query, groupID, receiverID).Scan(&status,&invitationID,)
	if err != nil && err.Error() != "no rows in result set" {
		log.Printf("database: Failed check for invitation: %v", err)
		return 0, err // Return error if failed to insert post
	}
	if status == "pending" {
		log.Printf("invite already made")
		return invitationID, nil
	}
	query = `
    INSERT INTO 
        group_invitations (group_id, sender_id, receiver_id) 
    VALUES 
        ($1, $2, $3)
		RETURNING
				invitation_id`
	err = DB.QueryRow(context.Background(), query, groupID, senderID, receiverID).Scan(&invitationID)
	if err != nil {
		log.Printf("database: Failed to insert invitation into database: %v", err)
		return 0, err // Return error if failed to insert post
	}
	return invitationID, nil
}

func RespondToInvite(response models.GroupResponse, userID string) error{
	query := `UPDATE group_invitations SET status = $1 WHERE group_id = $2 AND receiver_id = $3 AND status = 'pending'`
	_, err := DB.Exec(context.Background(), query, response.Status, response.GroupID, userID)
	if err != nil {
		log.Printf("database: Failed to update response in database: %v", err)
		return err // Return error if failed to update response
	}
	if response.Status == "accepted" {
		isMember, _ := GroupMember(userID, response.GroupID)
		if isMember{
			log.Printf("User already a member")
			return nil
		}
		query = `INSERT INTO 
			group_member (user_id, group_id)
	VALUES 
			($1, $2)`
		_, err = DB.Exec(context.Background(), query, userID, response.GroupID)
		if err != nil {
			log.Printf("database: Failed to add group member: %v", err)
			return err // Return error if failed to insert group member
		}
	} 
	return nil
}

func CreateRequest(groupID int, senderID string) (int, error) {
	var status string
	var requestID int
	var creatorID string
	query := `
	SELECT 
		 status,
		 request_id
		FROM
			group_requests
		WHERE
		 requester_id = $1 
`
	err := DB.QueryRow(context.Background(), query, senderID).Scan(&status,&requestID)
	if err != nil && err.Error() != "no rows in result set" {
		log.Printf("database: Failed check for request: %v", err)
		return 0, err // Return error if failed to insert post
	}
	if status == "pending" {
		log.Printf("request already made")
		return requestID, nil
	}
	query = `
    INSERT INTO 
        group_requests (group_id, requester_id) 
    VALUES 
        ($1, $2)
				RETURNING
				request_id`
	err = DB.QueryRow(context.Background(), query, groupID, senderID).Scan(&requestID)
	if err != nil {
		log.Printf("database: Failed to insert request into database: %v", err)
		return 0, err // Return error if failed to insert post
	}
	query = `SELECT creator_id FROM "group" WHERE group_id = $1`
	err = DB.QueryRow(context.Background(), query, groupID).Scan(&creatorID)
	if err != nil {
		log.Printf("database: Failed to insert request into database: %v", err)
		return 0, err // Return error if failed to insert post
	}
	// add to notification table for creator
	return requestID, nil
}

func RespondToRequest(response models.GroupResponse) error{
	query := `UPDATE group_requests SET status = $1 WHERE group_id = $2 AND requester_id = $3 AND status = 'pending'`
	_, err := DB.Exec(context.Background(), query, response.Status, response.GroupID, response.RequesterID)
	if err != nil {
		log.Printf("database: Failed to update response in database: %v", err)
		return err // Return error if failed to insert post
	}
	if response.Status == "accepted" {
		query = `INSERT INTO 
			group_member (user_id, group_id) 
	VALUES 
			($1, $2)`
		_, err = DB.Exec(context.Background(), query, response.RequesterID, response.GroupID)
		if err != nil {
			log.Printf("database: Failed to add group member: %v", err)
			return err // Return error if failed to insert post
		}
	} 
	return nil
}
