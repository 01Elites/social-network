package querys

import (
	"context"
	"errors"
	"log"

	"social-network/internal/models"
)

// CreatePostInDB adds the post to the database
func CreatePostInDB(userID string, post models.Create_Post) (int, error) {
	var postID int
	if post.GroupID != 0 { // post is a group post
		query := `
    INSERT INTO 
        post (title, content, privacy_type, group_id, user_id, image) 
    VALUES 
        ($1, $2, $3, $4, $5, $6)
    RETURNING post_id`
		err := DB.QueryRow(context.Background(), query, post.Title, post.Content, post.Privacy, post.GroupID, userID, post.Image).Scan(&postID)
		if err != nil {
			log.Printf("database: Failed to insert post into database: %v", err)
			return 0, err // Return error if failed to insert post
		}
	} else { // post is a user post
		query := `
    INSERT INTO 
        post (title, content, privacy_type, user_id, image) 
    VALUES 
        ($1, $2, $3, $4, $5)
    RETURNING post_id`
		err := DB.QueryRow(context.Background(), query, post.Title, post.Content, post.Privacy, userID, post.Image).Scan(&postID)
		if err != nil {
			log.Printf("database: Failed to insert post into database: %v", err)
			return 0, err // Return error if failed to insert post
		}
		if post.Privacy == "almost_private" { // add the userids for the users who are allowed to view post
			for i := 0; i < len(post.UserNames); i++ {
				ID, err := GetUserIDByUserName(post.UserNames[i])
				if err != nil {
					log.Printf("database: Failed to get user id: %v", err)
					return 0, err
				}
				query := `
		INSERT INTO
			post_user (post_id, allowed_user_id)
			VALUES
					($1, $2)`
				_, err = DB.Query(context.Background(), query, postID, ID)
				if err != nil {
					log.Printf("database: Failed to insert allowed user into database: %v", err)
					return 0, err // Return error if failed to insert post
				}
			}
		}
	}
	return postID, nil
}

// GetPostsFeed gets all the posts that the logged user can view
func GetPostsFeed(loggeduser models.User) ([]models.Post, error) {
	// Query the database
	query := `
    SELECT 
        post_id, 
        title,
				content, 
				created_at,
        user_id,
				user_name,
				post.image, 
        first_name,
				last_name,
				profile.image,
				privacy_type
    FROM 
        post 
    INNER JOIN 
        profile USING (user_id)
		INNER JOIN
				"user" USING (user_id)
		ORDER BY
				created_at DESC
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
			&p.User.UserName,
			&p.Image,
			&p.User.FirstName,
			&p.User.LastName,
			&p.User.Avatar,
			&p.PostPrivacy,
		); err != nil {
			log.Printf("database failed to scan post: %v\n", err)
			return nil, err
		}
		p.CommentsCount, err = GetCommentsCountByID(p.ID)
		if err != nil {
			log.Printf("database failed to scan comments count: %v\n", err)
			return nil, err
		}

		p.Likers_Usernames, p.IsLiked, err = GetPostLikers(p.ID, loggeduser.UserID)
		if err != nil {
			log.Printf("database: Failed to scan likers: %v\n", err)
			return []models.Post{}, err
		}
		p.PostLikes = len(p.Likers_Usernames)
		if p.PostPrivacy == "almost_private" {
			isAllowed, err := IsAllowed_AlmostPrivate(p.ID, loggeduser.UserID)
			if err != nil {
				log.Printf("database failed to scan allowed users: %v\n", err)
				return nil, err
			}
			if isAllowed {
				p.User.UserID = ""
				posts = append(posts, p)
			}
		} else if loggeduser.Following[p.User.UserID] || loggeduser.UserID == p.User.UserID {
			p.User.UserID = ""
			posts = append(posts, p)
		}
	}
	return posts, nil
}

func IsAllowed_AlmostPrivate(postID int, userID string) (bool, error) {
	query := `Select allowed_user_id FROM post_user WHERE post_id = $1`
	rows, err := DB.Query(context.Background(), query, postID)
	if err != nil {
		log.Printf("database failed to scan allowed users: %v\n", err)
		return false, err
	}
	for rows.Next() {
		var allowed_userid string
		if err := rows.Scan(&allowed_userid); err != nil {
			log.Printf("database failed to scan allowed_user: %v\n", err)
			return false, err
		}
		if allowed_userid == userID {
			return true, nil
		}
	}
	return false, nil
}

// Get post by ID gets the data for one post
func GetPostByID(postID int, userid string) (models.Post, error) {
	// Query the database
	query := `
    SELECT 
				post_id,
        title,
				content, 
				post.image,
        user_id, 
				user_name,
				created_at,
        first_name,
				last_name,
				profile.image
    FROM 
        post 
    INNER JOIN 
        profile USING (user_id)
		INNER JOIN
				"user" USING (user_id)
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
		&post.Image,
		&post.User.UserID,
		&post.User.UserName,
		&post.CreationDate,
		&post.User.FirstName,
		&post.User.LastName,
		&post.User.Avatar,
	)
	if err != nil {
		log.Printf("database: Failed to scan row: %v", err)
		return models.Post{}, err
	}
	post.CommentsCount, err = GetCommentsCountByID(post.ID)
	if err != nil {
		log.Printf("database failed to scan comments count: %v\n", err)
		return models.Post{}, err
	}

	post.Likers_Usernames, post.IsLiked, err = GetPostLikers(post.ID, userid)
	if err != nil {
		log.Printf("database: Failed to scan likers: %v\n", err)
		return models.Post{}, err
	}
	post.PostLikes = len(post.Likers_Usernames)
	post.User.UserID = ""
	return post, nil
}

// PostExistsByID checks if a post with the given postID exists
func PostExists(postID int) (bool, error) {
	query := "SELECT COUNT(*) FROM post WHERE post_id = $1"
	var count int
	if err := DB.QueryRow(context.Background(), query, postID).Scan(&count); err != nil {
		log.Printf("database: Failed to get number of posts: %v\n", err)
		return false, err
	}
	return count > 0, nil
}

// GetPostLikeCountByID gets the like counts for an indivitual post
func GetPostLikeCountByID(postID int) (int, error) {
	query := "SELECT COUNT(*) FROM post_interaction WHERE post_id = $1"
	var count int
	if err := DB.QueryRow(context.Background(), query, postID).Scan(&count); err != nil {
		log.Printf("database: Failed to num of likes: %v\n", err)
		return 0, err
	}
	return count, nil
}

// GetCommentsCountByID counts the number of comments for an indivitual post
func GetCommentsCountByID(postID int) (int, error) {
	query := "SELECT COUNT(*) FROM comment WHERE post_id = $1"
	var count int
	if err := DB.QueryRow(context.Background(), query, postID).Scan(&count); err != nil {
		log.Printf("database: Failed to get number of comments: %v\n", err)
		return 0, err
	}
	return count, nil
}

// GetPostLikers gets the usernames for all of the users who liked the post
func GetPostLikers(postID int, userID string) ([]string, bool, error) {
	var likers []string
	var isLiked bool
	query := `SELECT
	 user_name
	 FROM post_interaction
	 INNER JOIN
	 "user" USING (user_id)
	  WHERE post_id = $1`
	rows, err := DB.Query(context.Background(), query, postID)
	if err != nil {
		log.Printf("database: Failed to scan likers: %v\n", err)
		return nil, isLiked, err
	}
	for rows.Next() {
		var liker string
		err := rows.Scan(
			&liker,
		)
		if liker == userID {
			isLiked = true
		}
		if err != nil {
			log.Printf("database: Failed to scan liker: %v\n", err)
			return nil, isLiked, err
		}
		likers = append(likers, liker)
	}
	return likers, isLiked, nil
}

// DeletePost deletes the post from the database
func DeletePost(postID int, userID string) error {
	var creator string
	query := `SELECT user_id FROM post WHERE post_id = $1`
	err := DB.QueryRow(context.Background(), query, postID).Scan(&creator)
	if err != nil {
		log.Printf("Failed to get post creator: %v\n", err)
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

// GetGroupPosts selects all the posts for a single group
func GetGroupPosts(groupID int) ([]models.Post, error) {
	query := `
	SELECT 
			post_id, 
			title,
			content, 
			post.image,
			user_id, 
			user_name,
			profile.image
	FROM 
			post 
	INNER JOIN 
			profile USING (user_id)
	INNER JOIN
	 		"user" USING (user_id)
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
			&p.Image,
			&p.User.UserID,
			&p.User.UserName,
			&p.User.Avatar,
		); err != nil {
			log.Printf("database failed to scan post: %v\n", err)
			return nil, err
		}

		p.CommentsCount, err = GetCommentsCountByID(p.ID)
		if err != nil {
			log.Printf("database failed to scan comments count: %v\n", err)
			return nil, err
		}
		p.Likers_Usernames, p.IsLiked, err = GetPostLikers(p.ID, "")
		if err != nil {
			log.Printf("database: Failed to scan likers: %v\n", err)
			return []models.Post{}, err
		}
		p.User.UserID = ""
		posts = append(posts, p)
	}
	return posts, nil
}

// GetUserPosts gets the posts created by a single user
func GetUserPosts(loggeduser string, userid string, followed bool) ([]models.Post, error) {
	user, err := GetUserProfile(userid)
	if err != nil {
		log.Printf("database: Failed to get user profile: %v\n", err)
		return nil, err
	}
	// Query the database
	query := `
    SELECT 
				post_id,
        title,
				content, 
        privacy_type,
				created_at,
				image
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
	posts := make([]models.Post, 0)
	// Iterate through the rows
	for rows.Next() {
		post := models.Post{
			User: models.PostFeedProfile{
				UserName:  user.Username,
				FirstName: user.FirstName,
				LastName:  user.LastName,
				Avatar:    user.Avatar,
			},
		}
		if err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.PostPrivacy,
			&post.CreationDate,
			&post.Image,
		); err != nil {
			log.Printf("database failed to scan post: %v\n", err)
			return nil, err
		}
		post.CommentsCount, err = GetCommentsCountByID(post.ID)
		if err != nil {
			log.Printf("database failed to scan comments count: %v\n", err)
			return nil, err
		}
		post.Likers_Usernames, post.IsLiked, err = GetPostLikers(post.ID, loggeduser)
		if err != nil {
			log.Printf("database: Failed to scan likers: %v\n", err)
			return []models.Post{}, err
		}
		post.PostLikes = len(post.Likers_Usernames)

		if userid == loggeduser {
			posts = append(posts, post)
			continue
		}
		if post.PostPrivacy == "public" {
			posts = append(posts, post)
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
