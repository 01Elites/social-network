package database

import (
	"context"
	"log"

	"social-network/internal/models"
)

func GetGroupPosts(groupID int) ([]models.Post, error) {
	query := `
	SELECT 
			post_id, 
			title,
			content, 
			user_id, 
			nick_name,
			privacy_type
	FROM 
			post 
	INNER JOIN 
			profile USING (user_id)
	WHERE
			group_id = $1`
	rows, err := DB.Query(context.Background(), query, groupID)
	if err != nil {
		log.Printf("database failed to scan post: %v\n", err)
		return nil, err
	}
	defer rows.Close()
	var posts []models.Post
	for rows.Next() {
		var p models.Post
		if err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Content,
			&p.User.UserID,
			&p.User.NickName,
			&p.PostPrivacy,
		); err != nil {
			log.Printf("database failed to scan post: %v\n", err)
			return nil, err
		}
		posts = append(posts, p)
	}
	return posts, nil
}
