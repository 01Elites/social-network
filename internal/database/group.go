package database

import (
	"context"
	"log"
	"social-network/internal/events"
)

func CreateGroup(userID string, group events.Create_Group) error{
	query := `
    INSERT INTO 
        "group" (title, description, creator_id) 
    VALUES 
        ($1, $2, $3)
		RETURNING group_id`
		var group_id int
	err := DB.QueryRow(context.Background(),query, group.Title, group.Description, userID).Scan(&group_id)
	if err != nil {
		log.Printf("database: Failed to insert group into database: %v", err)
		return err // Return error if failed to insert post
	}
	return nil
}