package database

import ("social-network/internal/events"
				"log"
			"context")

func CreatePostInDB(userID string, post events.Create_Post) error{
	query := `
    INSERT INTO 
        post (title, post, user_id, privacy_type) 
    VALUES 
        (?, ?, ?, ?)
    RETURNING p_id`
		var post_id int // Assuming p_id is of type int
	err := DB.QueryRow(context.Background(),query, post.Title, post.Content, post.Privacy, userID).Scan(&post_id)
	if err != nil {
		log.Printf("database: Failed to insert post into database: %v", err)
		return err // Return error if failed to insert post
	}
	return nil
}

func GetPostsFeed(userID, page, catID int) ([]events.PostFeed, error) {
	// Query the database
	query := `
    SELECT 
        post.post_id, 
        post.title,
				post.content, 
        post.creation_date,
				post.group_id, 
        post.user_id, 
        user.user_name, 
        user.first_name, 
        user.last_name, 
        user.gender,
    FROM 
        post 
    INNER JOIN 
        user ON post.user_id = user.user_id
		IF post.privacy_type = 'group'
		INNER JOIN
				group_member ON group_member.user_id = ? AND group_member.group_id = post.group_id
		INNER JOIN
				group ON group.group_id = post.group_id
		ELSE
		BEGIN
		IF post.privacy_type = 'almost_private'
		INNER JOIN 
			post_user ON post_user.post_id = post.post_id AND post_user.allowed_user_id = ?
		ELSE
			INNER JOIN 
						follower ON follower.followed_id = ? OR follower.follower_id = ?
		END`

	query += `
	GROUP BY
		post.p_id
	ORDER BY 
		post.p_id DESC`

	rows, err := DB.Query(context.Background(), query, userID, userID, userID, userID)
	if err != nil {
		log.Printf("database failed to scan post: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	// Create a slice to hold the results
	posts := make([]events.PostFeed, 0)
	// Iterate through the rows
	for rows.Next() {
		var p events.PostFeed
		if err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Content,
			&p.CreationDate,
			&p.GroupID,
			&p.User.ID,
			&p.User.UserName,
			&p.User.FirstName,
			&p.User.LastName,
			&p.User.Gender,
		); err != nil {
			log.Printf("database failed to scan post: %v\n", err)
			return nil, err
		}
		posts = append(posts, p)
	}

	return posts, nil
}