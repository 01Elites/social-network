package database

import (
	"context"
	"fmt"
	"log"
	"social-network/internal/events"
)

func CreatePostInDB(userID string, post events.Create_Post) error{
	query := `
    INSERT INTO 
        post (title, content, privacy_type, group_id, user_id) 
    VALUES 
        ($1, $2, $3, $4, $5)
    RETURNING post_id`
		var post_id int // Assuming p_id is of type int
	err := DB.QueryRow(context.Background(),query, post.Title, post.Content, post.Privacy, post.GroupID, userID).Scan(&post_id)
	fmt.Println(post_id)
	if err != nil {
		log.Printf("database: Failed to insert post into database: %v", err)
		return err // Return error if failed to insert post
	}
	return nil
}

func GetPostsFeed(loggeduser events.User) ([]events.Post, error) {
	// Query the database
	query := `
    SELECT 
        post_id, 
        title,
				content, 
        user_id, 
        user_name,
				privacy_type,
				group_id
    FROM 
        post 
    INNER JOIN 
        "user" USING (user_id)`


	rows, err := DB.Query(context.Background(), query)
	if err != nil {
		log.Printf("database failed to scan post: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	// Create a slice to hold the results
	posts := make([]events.Post, 0)
	// Iterate through the rows
	for rows.Next() {
		var p events.Post
		if err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Content,
			&p.User.ID,
			&p.User.UserName,
			&p.PostPrivacy,
			&p.GroupID,
		); err != nil {
			log.Printf("database failed to scan post: %v\n", err)
			return nil, err
		}
		if p.PostPrivacy == "group" && loggeduser.Groups[p.GroupID]{
			posts = append(posts, p)
		} else if p.PostPrivacy == "almost_private" {
			query = `Select allowed_user_id FROM post_user WHERE post_id = $1`
			rows, err := DB.Query(context.Background(), query)
	if err != nil {
		log.Printf("database failed to scan allowed users: %v\n", err)
		return nil, err
	}
		for rows.Next(){
			var allowed_userid string 
			if err := rows.Scan(&allowed_userid);err != nil {
				log.Printf("database failed to scan allowed_user: %v\n", err)
			return nil, err
			}
			if allowed_userid == loggeduser.ID {
				posts = append(posts, p)
				break
			}
		}
		} else if loggeduser.Following[p.User.ID]{
			posts = append(posts, p)
	} 
	}
	return posts, nil
}


func GetPostByID(postID int) (events.PostFeed, error) {
	// Query the database
	query := `
    SELECT 
         post_id, 
        title,
				content, 
        user_id, 
        user_name,
    FROM 
        post 
    INNER JOIN 
        user USING (user_id)
    WHERE 
        post.p_id = ? AND post.comment_of = 0
`

	// Execute the query and retrieve the row
	row := DB.QueryRow(context.Background(), query, postID)

	// Create a Post object to hold the result
	var post events.PostFeed

	// Scan the row into the Post object
	err := row.Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.User.ID,
		&post.User.UserName,
	)
	if err != nil {
		log.Printf("database: Failed to scan row: %v", err)
		return events.PostFeed{}, err
	}
	return post, nil
}
