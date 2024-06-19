package database

import ("social-network/internal/events"
				"log"
			"context"
			"fmt")

func CreatePostInDB(userID int, post events.Create_Post) error{
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
        post.p_id, 
        post.title, 
        post.creation_date, 
        post.user_id, 
        user.user_name, 
        user.first_name, 
        user.last_name, 
        user.gender,
    FROM 
        post 
    INNER JOIN 
        user ON post.user_id = user.user_id
    INNER JOIN 
        follower ON post.user_id = follower.followed_id AND user.user_id = follower.follower_id
    LEFT JOIN 
        category ON threads.cat_id = category.cat_id
    LEFT JOIN 
        posts_interaction AS likes ON post.p_id = likes.post_id AND likes.user_id = ?
    WHERE 
        post.comment_of = 0 `

	// Conditionally add WHERE clause to filter by category
	if catID != 0 {
		query += fmt.Sprintf("AND category.cat_id = %d", catID)
	}

	query += `
	GROUP BY
		post.p_id
	ORDER BY 
		post.p_id DESC`

	rows, err := DB.Query(query, userID, page*10)
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
			&p.CreationDate,
			&p.User.ID,
			&p.User.UserName,
			&p.User.FirstName,
			&p.User.LastName,
			&p.User.Gender,
			&p.Category,
			&p.PostLikes,
			&p.IsLiked,
			&p.CommentsCount,
			&p.ViewsCount,
		); err != nil {
			log.Printf("database failed to scan post: %v\n", err)
			return nil, err
		}
		posts = append(posts, p)
	}

	return posts, nil
}