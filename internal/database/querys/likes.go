package querys

import (
	"context"
	"log"
)

// Uptate_Like_in_db updates the like in the database.
// for both posts and comments
func UpDateLikeInDB(userID string, PostID int) error {
	// Check if the user has already liked the post
	query := "SELECT COUNT(*) FROM post_interaction WHERE user_id = $1 AND post_id = $2"
	var count int
	if err := DB.QueryRow(context.Background(), query, userID, PostID).Scan(&count); err != nil {
		log.Printf("database: Failed to num of posts: %v\n", err)
		return err
	}

	// Check if the user has already liked the post
	var final_query string
	if count > 0 { // If the user has already liked the post
		query := "DELETE FROM post_interaction WHERE user_id = $1 AND post_id = $2"
		_, err := DB.Exec(context.Background(), query, userID, PostID)
		if err != nil {
			log.Printf("database: Failed to delete like: %v\n", err)
			return err
		}
	} else {
		final_query = "INSERT INTO post_interaction (user_id, post_id, interaction_type) VALUES ($1, $2, $3)"
		// Update or insert the like
		if _, err := DB.Exec(context.Background(), final_query, userID, PostID, "like"); err != nil {
			log.Printf("database: Failed to update like: %v\n", err)
			return err
		}
	}
	return nil
}
