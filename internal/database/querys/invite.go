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
		log.Printf("request already made")
		return invitationID, nil
	}
	query = `
    INSERT INTO 
        group_invitations (group_id, sender_id, reciever_id) 
    VALUES 
        ($1, $2, $3)`
	err = DB.QueryRow(context.Background(), query, groupID, senderID, receiverID).Scan(&invitationID)
	if err != nil {
		log.Printf("database: Failed to insert post into database: %v", err)
		return 0, err // Return error if failed to insert post
	}
	return invitationID, nil
}

func RespondToInvite(response models.InviteResponse, userID string) error{
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