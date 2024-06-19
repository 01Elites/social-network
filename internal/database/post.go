package database

import ("social-network/internal/events"
				"log"
			"context")

func CreatePostInDB(userID int, post events.Create_Post) error{
	query := `
    INSERT INTO 
        post (title, post, user_id) 
    VALUES 
        (?, ?, ?)
    RETURNING p_id`
		var post_id int // Assuming p_id is of type int
	err := DB.QueryRow(context.Background(),query, post.Title, post.Content, userID).Scan(&post_id)
	if err != nil {
		log.Printf("database: Failed to insert post into database: %v", err)
		return err // Return error if failed to insert post
	}
	return nil
}