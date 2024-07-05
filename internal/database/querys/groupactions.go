package querys

import (
	"context"
	"log"
	"social-network/internal/models"
)

func CreateInvite(groupID int, senderID string, receiverID string) (int, error) {
	var status string
	var invitationID int
	query := `
	SELECT 
		 status,
		 invitation_id
		FROM
			group_invitation
		WHERE
		 (sender_id = $1 AND receiver_id = $2) 
`
	err := DB.QueryRow(context.Background(), query, senderID, receiverID).Scan(invitationID, status)
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
        group_invitations (group_id, sender_id, reciever_id) 
    VALUES 
        ($1, $2, $3)`
	err = DB.QueryRow(context.Background(), query, groupID, senderID, receiverID).Scan(&invitationID)
	if err != nil {
		log.Printf("database: Failed to insert invitation into database: %v", err)
		return 0, err // Return error if failed to insert post
	}
	return invitationID, nil
}

func RespondToInvite(response models.GroupResponse, userID string) error{
	if response.Status == "accepted" {
		query := `UPDATE group_invitations SET status = $1 WHERE group_id = $2, receiver_id = $3 AND status = 'pending'`
		_, err := DB.Exec(context.Background(), query, "accepted", response.GroupID, userID)
		if err != nil {
			log.Printf("database: Failed to update response in database: %v", err)
			return err // Return error if failed to insert post
		}
		if response.ID == 0 {
			log.Print("no match")
			return nil
		}
		query = `INSERT INTO 
			group_member (user_id, group_id) 
	VALUES 
			($1, $2)`
		_, err = DB.Exec(context.Background(), query, userID, response.GroupID)
		if err != nil {
			log.Printf("database: Failed to add follower: %v", err)
			return err // Return error if failed to insert post
		}
	} else if response.Status == "rejected" {
		query := `UPDATE group_invitations SET status = 'rejected' WHERE sender_id = $1 AND group_id = $2`
		_, err := DB.Exec(context.Background(), query, userID, response.GroupID)
		if err != nil {
			log.Printf("database: Failed to update response in database: %v", err)
			return err // Return error if failed to insert post
		}
	}
	return nil
}

func CreateRequest(groupID int, senderID string) (int, error) {
	var status string
	var requestID int
	query := `
	SELECT 
		 status,
		 request_id
		FROM
			group_requests
		WHERE
		 sender_id = $1 
`
	err := DB.QueryRow(context.Background(), query, senderID).Scan(requestID, status)
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
        group_requests (group_id, sender_id) 
    VALUES 
        ($1, $2)`
	err = DB.QueryRow(context.Background(), query, groupID, senderID).Scan(&requestID)
	if err != nil {
		log.Printf("database: Failed to insert request into database: %v", err)
		return 0, err // Return error if failed to insert post
	}
	return requestID, nil
}

func RespondToRequest(response models.GroupResponse) error{
	var requesterID string
	if response.Status == "accepted" {
		query := `UPDATE group_requests SET status = $1 WHERE group_id = $2, receiver_id = $3 AND status = 'pending'`
		_, err := DB.Exec(context.Background(), query, "accepted", response.GroupID, response.RequesterID)
		if err != nil {
			log.Printf("database: Failed to update response in database: %v", err)
			return err // Return error if failed to insert post
		}
		if response.ID == 0 {
			log.Print("no match")
			return nil
		}
		query = `INSERT INTO 
			group_member (user_id, group_id) 
	VALUES 
			($1, $2)`
		_, err = DB.Exec(context.Background(), query, requesterID, response.GroupID)
		if err != nil {
			log.Printf("database: Failed to add group member: %v", err)
			return err // Return error if failed to insert post
		}
	} else if response.Status == "rejected" {
		query := `UPDATE group_requests SET status = 'rejected' WHERE requester_id = $1 AND group_id = $2`
		_, err := DB.Exec(context.Background(), query, response.RequesterID, response.GroupID)
		if err != nil {
			log.Printf("database: Failed to update response in database: %v", err)
			return err // Return error if failed to insert post
		}
	}
	return nil
}