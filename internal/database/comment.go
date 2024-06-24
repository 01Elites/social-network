package database

import (
	"context"
	"log"
	"social-network/internal/events"
)

func Create_Comment_in_db(userID string, comment events.Create_Comment) error {
	// Insert the comment into the database
	query := `
    INSERT INTO 
        comment (post_id, user_id , content) 
    VALUES 
        ($1, $2, $3)`

	_, err := DB.Exec(context.Background(), query, comment.ParentID,userID, comment.Content)
	if err != nil {
		log.Printf("database: Failed to insert comment into database: %v", err)
		return err
	}

	return nil
}

func Get_PostComments_from_db(userID string, postID, page int) ([]events.Comment, error) {
	// Query the database
	query := `
    SELECT 
        comment_id, 
				content, 
        user_id, 
        user_name, 
    FROM 
        comment 
    INNER JOIN 
        user USING (user_id)
    WHERE 
        post_id = ?`

	// Execute the query and retrieve the rows
	rows, err := DB.Query(context.Background(), query, postID)
	if err != nil {
		log.Printf("database failed to scan post: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	// Create a slice to hold the results
	var comments []events.Comment

	// Iterate through the rows
	for rows.Next() {
		var comment events.Comment
		if err := rows.Scan(
			&comment.ID,
			&comment.Content,
			&comment.User.ID,
			&comment.User.UserName,
		); err != nil {
			log.Printf("database failed to scan post: %v\n", err)
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}