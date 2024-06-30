package database

import (
	"context"
	"log"
)

func CreateInvite(groupID int, from string, to string) error {
	query := `
    INSERT INTO 
        group_invitations (group_id, sender_id, reciever_id) 
    VALUES 
        ($1, $2, $3)`
	_, err := DB.Query(context.Background(), query, groupID, from, to)
	if err != nil {
		log.Printf("database: Failed to insert post into database: %v", err)
		return err // Return error if failed to insert post
	}
	return nil
}
