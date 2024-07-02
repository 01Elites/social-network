package database

import (
	"context"
	"log"
	"social-network/internal/models"
	"errors"
)

func CreatePostInDB(userID string, post models.Create_Post) (int, error){
	var postID int
	if post.GroupID != 0 {
	query := `
    INSERT INTO 
        post (title, content, privacy_type, group_id, user_id) 
    VALUES 
        ($1, $2, $3, $4, $5)
    RETURNING post_id`
	err := DB.QueryRow(context.Background(),query, post.Title, post.Content, post.Privacy, post.GroupID, userID).Scan(&postID)
	if err != nil {
		log.Printf("database: Failed to insert post into database: %v", err)
		return 0, err // Return error if failed to insert post
	}
} else {
	query := `
    INSERT INTO 
        post (title, content, privacy_type, user_id) 
    VALUES 
        ($1, $2, $3, $4)
    RETURNING post_id`
	err := DB.QueryRow(context.Background(),query, post.Title, post.Content, post.Privacy, userID).Scan(&postID)
	if err != nil {
		log.Printf("database: Failed to insert post into database: %v", err)
		return 0, err // Return error if failed to insert post
	}
	if post.Privacy == "almost_private" {
		for i:=0;i<len(post.UserIDs);i++{
		query := `
		INSERT INTO
			post_user (post_id, allowed_user_id)
			VALUES
					($1, $2)`
		_, err := DB.Query(context.Background(), query, postID, post.UserIDs[i])
		if err != nil {
			log.Printf("database: Failed to insert post into database: %v", err)
			return 0, err // Return error if failed to insert post
		}
	}
	}
}
	return postID, nil
}

func GetPostsFeed(loggeduser models.User) ([]models.Post, error) {
	// Query the database
	query := `
    SELECT 
        post_id, 
        title,
				content, 
				created_at,
        user_id, 
        first_name,
				last_name,
				privacy_type
    FROM 
        post 
    INNER JOIN 
        profile USING (user_id)
`
	rows, err := DB.Query(context.Background(), query)
	if err != nil {
		log.Printf("database failed to scan post: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	// Create a slice to hold the results
	posts := make([]models.Post, 0)
	// Iterate through the rows
	for rows.Next() {
		var p models.Post
		if err := rows.Scan(
			&p.ID,
			&p.Title,
			&p.Content,
			&p.CreationDate,
			&p.User.UserID,
			&p.User.FirstName,
			&p.User.LastName,
			&p.PostPrivacy,
		); err != nil {
			log.Printf("database failed to scan post: %v\n", err)
			return nil, err
		}
		p.CommentsCount, err = GetCommentsCountByID(p.ID)
		if err != nil{
			log.Printf("database failed to scan comments count: %v\n", err)
			return nil, err
		}
		p.Likers_ids,p.IsLiked, err = GetPostLikers(p.ID, loggeduser.UserID)
		if err != nil {
			log.Printf("database: Failed to scan likers: %v\n", err)
			return []models.Post{}, err
		}
		p.PostLikes = len(p.Likers_ids)
	 if p.PostPrivacy == "almost_private" {
			query = `Select allowed_user_id FROM post_user WHERE post_id = $1`
			rows, err := DB.Query(context.Background(), query, p.ID)
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
			if allowed_userid == loggeduser.UserID {
				posts = append(posts, p)
				break
			}
		}
		} else if loggeduser.Following[p.User.UserID]{
			posts = append(posts, p)
	} 
	}
	return posts, nil
}


func GetPostByID(postID int, userid string) (models.Post, error) {
	// Query the database
	query := `
    SELECT 
				post_id,
        title,
				content, 
        user_id, 
				created_at,
        first_name,
				last_name
    FROM 
        post 
    INNER JOIN 
        profile USING (user_id)
    WHERE 
        post_id = $1
`

	// Execute the query and retrieve the row
	row := DB.QueryRow(context.Background(), query, postID)

	// Create a Post object to hold the result
	var post models.Post
	// Scan the row into the Post object
	err := row.Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.User.UserID,
		&post.CreationDate,
		// &post.User.Image,
		&post.User.FirstName,
		&post.User.LastName,
	)
	if err != nil {
		log.Printf("database: Failed to scan row: %v", err)
		return models.Post{}, err
	}
	post.CommentsCount, err = GetCommentsCountByID(post.ID)
	if err != nil{
		log.Printf("database failed to scan comments count: %v\n", err)
		return models.Post{}, err
	}
	post.Likers_ids, post.IsLiked, err = GetPostLikers(post.ID, userid)
	if err != nil {
		log.Printf("database: Failed to scan likers: %v\n", err)
		return models.Post{}, err
	}
	post.PostLikes = len(post.Likers_ids)
	return post, nil
}

func GetPostCountByID(postID int) (int, error) {
	query := "SELECT COUNT(*) FROM post WHERE post_id = $1"
	var count int
	if err := DB.QueryRow(context.Background(), query, postID).Scan(&count); err != nil {
		log.Printf("database: Failed to num of posts: %v\n", err)
		return 0, err
	}
	return count, nil
}

func GetPostLikeCountByID(postID int) (int, error) {
	query := "SELECT COUNT(*) FROM post_interaction WHERE post_id = $1"
	var count int
	if err := DB.QueryRow(context.Background(), query, postID).Scan(&count); err != nil {
		log.Printf("database: Failed to num of likes: %v\n", err)
		return 0, err
	}
	return count, nil
}

func GetCommentsCountByID(postID int)(int, error){
	query := "SELECT COUNT(*) FROM comment WHERE post_id = $1"
	var count int
	if err := DB.QueryRow(context.Background(), query, postID).Scan(&count); err != nil {
		log.Printf("database: Failed to get number of comments: %v\n", err)
		return 0, err
	}
	return count, nil
}

func GetPostLikers(postID int, userID string)([]string, bool, error){
	var likers []string
	var isLiked bool
	query := `SELECT user_id FROM post_interaction WHERE post_id = $1`
	rows, err := DB.Query(context.Background(), query, postID)
	if err != nil {
		log.Printf("database: Failed to scan likers: %v\n", err)
		return nil, isLiked, err
	}
	for rows.Next(){
		var liker string
		err := rows.Scan(
			&liker,
		)
		if liker == userID {
			isLiked = true
		}
		if err != nil{
			log.Printf("database: Failed to scan liker: %v\n", err)
		return nil, isLiked, err
		}
		likers = append(likers, liker)
	}
	return likers,isLiked, nil
}

func DeletePost(postID int, userID string) error {
	var creator string
	query := `SELECT user_id FROM post WHERE post_id = $1`
	err:= DB.QueryRow(context.Background(), query, postID).Scan(&creator)
	if err != nil {
		return err
	}
	if creator != userID {
		log.Print("user unauthorized")
		return errors.New("user unauthorized")
	}
	query = `DELETE FROM post WHERE post_id = $1`
	_, err = DB.Exec(context.Background(), query, postID)
	if err != nil {
		log.Printf("failed to delete POST:%v\n", err)
		return err
	}
	return nil
}

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

func GetUserPosts(loggeduser string, userid string, followed bool) ([]models.ProfilePost, error) {
	// Query the database
	query := `
    SELECT 
				post_id,
        title,
				content, 
        privacy_type,
				created_at
    FROM 
        post 
    WHERE 
        user_id = $1
`

	// Execute the query and retrieve the row
	rows, err := DB.Query(context.Background(), query, userid)
	if err != nil {
		log.Printf("database failed to scan post: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	// Create a slice to hold the results
	posts := make([]models.ProfilePost, 0)
	// Iterate through the rows
	for rows.Next() {
		var post models.ProfilePost
		if err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.PostPrivacy,
			&post.CreationDate,
		); err != nil {
			log.Printf("database failed to scan post: %v\n", err)
			return nil, err
		}
		post.CommentsCount, err = GetCommentsCountByID(post.ID)
		if err != nil{
			log.Printf("database failed to scan comments count: %v\n", err)
			return nil, err
		}
		post.Likers_ids,post.IsLiked, err = GetPostLikers(post.ID, loggeduser)
		if err != nil {
			log.Printf("database: Failed to scan likers: %v\n", err)
			return []models.ProfilePost{}, err
		}
		post.PostLikes = len(post.Likers_ids)
		if userid == loggeduser {
			posts = append(posts, post)
			continue
		}
		if post.PostPrivacy == "public"{
			posts= append(posts, post)
		} else if post.PostPrivacy == "almost_private" {
			var count int
			query = `SELECT COUNT(*) FROM post_user WHERE post_id = $1 AND allowed_user_id = $2`
			DB.QueryRow(context.Background(), query, post.ID, loggeduser).Scan(&count)
			if count > 0 {
				posts = append(posts, post)
			}
		} else if post.PostPrivacy == "private" {
			if followed {
				posts = append(posts, post)
			}
		}
	}
	return posts, nil

}