package querys

import (
	"context"
	"log"

	"social-network/internal/models"
)

// Create_Comment_in_db adds the comment to the database
func Create_Comment_in_db(userID string, comment models.Create_Comment) error {
	// Insert the comment into the database
	query := `
    INSERT INTO 
        comment (post_id, user_id , image, content) 
    VALUES 
        ($1, $2, $3, $4)`

	_, err := DB.Exec(context.Background(), query, comment.ParentID, userID, comment.Image, comment.Content)
	if err != nil {
		log.Printf("database: Failed to insert comment into database: %v", err)
		return err
	}
	return nil
}

// Get_PostComments_from_db gets all the comments pretaining to a post.
func Get_PostComments_from_db(userID string, postID int) ([]models.Comment, error) {
	// Query the database
	query := `
    SELECT 
        comment_id, 
				content, 
        user_id,
				comment.image,
				first_name,
				last_name,
				created_at,
				profile.image
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
			&comment.Image,
			&comment.User.FirstName,
			&comment.User.LastName,
			&comment.CreationDate,
			&comment.User.Avatar,
		); err != nil {
			log.Printf("database failed to scan post: %v\n", err)
			return nil, err
		}

		comment.User.UserName, err = GetUserNameByID(userID)
		if err != nil {
			log.Printf("database failed to get username: %v\n", err)
			return nil, err
		}
		comment.User.UserID = ""
		comments = append(comments, comment)
	}
	return comments, nil
}
