package database

import (
	"context"
	"log"
	"social-network/internal/models"
)

func Create_Comment_in_db(userID string, comment models.Create_Comment) error {
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

func Get_PostComments_from_db(userID string, postID, page int) ([]models.Comment, error) {
	// Query the database
	query := `
    SELECT 
        comment_id, 
				content, 
        user_id,
				first_name,
				last_name
    FROM 
        comment
		INNER JOIN 
        profile USING (user_id)
    WHERE 
        post_id = $1`

	// Execute the query and retrieve the rows
	rows, err := DB.Query(context.Background(), query, postID)
	if err != nil {
		log.Printf("database failed to scan post: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	// Create a slice to hold the results
	var comments []models.Comment

	// Iterate through the rows
	for rows.Next() {
		var comment models.Comment
		if err := rows.Scan(
			&comment.ID,
			&comment.Content,
			&comment.User.UserID,
			&comment.User.FirstName,
			&comment.User.LastName,
		); err != nil {
			log.Printf("database failed to scan post: %v\n", err)
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}