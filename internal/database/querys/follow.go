package querys

import (
	"context"
	"log"

	"social-network/internal/models"
)

func GetFollowingCount(userID string) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM follower WHERE follower_id = $1`
	err := DB.QueryRow(context.Background(), query, userID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func GetFollowerCount(userID string) (int, error) {
	var count int
	query := `SELECT COUNT(*) FROM follower WHERE followed_id = $1`
	err := DB.QueryRow(context.Background(), query, userID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func GetUsersFollowingByID(userID string) (map[string]bool, error) {
	Following := make(map[string]bool)
	query := `SELECT followed_id FROM follower WHERE follower_id = $1`
	rows, err := DB.Query(context.Background(), query, userID)
	if err != nil {
		log.Printf("database failed to scan followed user: %v\n", err)
		return nil, err
	}
	for rows.Next() {
		var followedUserID string
		if err = rows.Scan(&followedUserID); err != nil {
			log.Printf("database failed to scan followed user: %v\n", err)
			return nil, err
		}
		Following[followedUserID] = true
	}
	return Following, nil
}

func GetUserFollowingUserNames(userID string) ([]string, error) {
	var following []string
	query := `
	SELECT
    "user".user_name
	FROM
			follower
	INNER JOIN
			public."user" ON follower.followed_id = public."user".user_id
	WHERE
			follower.follower_id = $1
	`
	rows, err := DB.Query(context.Background(), query, userID)
	if err != nil {
		log.Printf("database failed to scan followed user: %v\n", err)
		return nil, err
	}
	for rows.Next() {
		var followedUsername string
		if err = rows.Scan(&followedUsername); err != nil {
			log.Printf("database failed to scan followed user: %v\n", err)
			return nil, err
		}
		following = append(following, followedUsername)
	}
	return following, nil
}

func GetUserFollowerUserNames(userID string) ([]string, error) {
	var followers []string
	query := `
	SELECT
		"user".user_name
	FROM
		follower
	INNER JOIN
		public."user" ON follower.follower_id = public."user".user_id
	WHERE
		follower.followed_id = $1
	`
	rows, err := DB.Query(context.Background(), query, userID)
	if err != nil {
		log.Printf("database failed to scan followers: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var followerUsername string
		if err = rows.Scan(&followerUsername); err != nil {
			log.Printf("database failed to scan follower user: %v\n", err)
			return nil, err
		}
		followers = append(followers, followerUsername)
	}

	if err = rows.Err(); err != nil {
		log.Printf("rows iteration error: %v\n", err)
		return nil, err
	}

	return followers, nil
}

func GetUsersFollowees(userID string) (map[string]bool, error) {
	Followees := make(map[string]bool)
	query := `SELECT follower_id FROM follower WHERE followed_id = $1`
	rows, err := DB.Query(context.Background(), query, userID)
	if err != nil {
		log.Printf("database failed to scan follower user: %v\n", err)
		return nil, err
	}
	for rows.Next() {
		var followerUserID string
		if err = rows.Scan(&followerUserID); err != nil {
			log.Printf("database failed to scan follower user: %v\n", err)
			return nil, err
		}
		Followees[followerUserID] = true
	}
	return Followees, nil
}

func CreateFollowRequest(request *models.Request) error {
	// Check if request already exists
	query := `
		SELECT 
				request_id, status
		FROM
				follow_requests
		WHERE
				sender_id = $1 
				AND receiver_id = $2
				AND status NOT IN ('accepted', 'rejected' , 'canceled');
`
	err := DB.QueryRow(context.Background(), query, request.Sender, request.Receiver).Scan(&request.ID, &request.Status)
	if err != nil && err.Error() != "no rows in result set" {
		log.Printf("database: Failed check for request: %v", err)
		return err // Return error if failed to insert post
	}

	// if it is already sent cancel the request by updating the status to canceled
	if request.Status == "pending" {
		query := `UPDATE follow_requests SET status = 'canceled' WHERE sender_id = $1 AND receiver_id = $2 AND request_id = $3`
		err := DB.QueryRow(context.Background(), query, request.Sender, request.Receiver, request.ID).Scan(&request.ID)
		if err != nil && err.Error() != "no rows in result set" {
			log.Printf("database1: Failed to update request in database: %v", err)
			return err
		}
		request.Status = "canceled"
		if err := CancelNotification(request.ID, "follow_request", request.Receiver); err != nil {
			log.Printf("database: Failed to cancel notification: %v", err)
			return err
		}
		return nil
	}

	// Insert request into database
	query = `
	INSERT INTO 
			follow_requests (sender_id, receiver_id) 
	VALUES 
			($1, $2)
	RETURNING request_id`
	err = DB.QueryRow(context.Background(), query, request.Sender, request.Receiver).Scan(&request.ID)
	if err != nil {
		log.Printf("database: Failed to insert request into database: %v", err)
		return err // Return error if failed to insert post
	}
	request.Status = "pending"
	return nil
}

func RespondToFollow(response models.Response) error {
	query := `UPDATE follow_requests SET status = $1 WHERE sender_id = $2 AND receiver_id = $3 RETURNING request_id`
	err := DB.QueryRow(context.Background(), query, response.Status, response.Follower, response.Followee).Scan(&response.ID)
	if err != nil {
		log.Printf("database: Failed to update response in database: %v", err)
		return err // Return error if failed to insert post
	}
	if response.ID == 0 {
		log.Print("no match")
		return nil
	}
	if response.Status == "accepted" {
		query = `INSERT INTO 
			follower (follower_id, followed_id) 
	VALUES 
			($1, $2)`
		_, err = DB.Exec(context.Background(), query, response.Follower, response.Followee)
		if err != nil {
			log.Printf("database: Failed to add follower: %v", err)
			return err // Return error if failed to insert post
		}
	}
	return nil
}

// GetFollowRequests returns all the follow requests for a user
func GetFollowRequests(userID string) ([]models.FriendRequest, error) {
	var requests []models.FriendRequest
	query := `
	SELECT
			"user".user_name,
			follow_requests.created_at
	FROM
			follow_requests
	INNER JOIN
			public."user" ON follow_requests.sender_id = public."user".user_id
	WHERE
      follow_requests.receiver_id = $1 AND follow_requests.status = 'pending'
	`

	rows, err := DB.Query(context.Background(), query, userID)
	if err != nil {
		log.Printf("Database query error: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var request models.FriendRequest
		if err := rows.Scan(&request.UserName, &request.Creation_date); err != nil {
			log.Printf("Row scan error: %v\n", err)
			return nil, err
		}
		requests = append(requests, request)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Rows iteration error: %v\n", err)
		return nil, err
	}
	return requests, nil
}

// GetExplore returns all the users that are not friends or requested to be friends
func GetExplore(userID string) ([]string, error) {
	var explore []string
	query := `
	SELECT
		"user".user_name
	FROM
		public."user"
	WHERE
		"user".user_id NOT IN (
			SELECT
				follower.followed_id
			FROM
				follower
			WHERE
				follower.follower_id = $1
		)
		AND "user".user_id NOT IN (
			SELECT
				follow_requests.sender_id
			FROM
				follow_requests
			WHERE
				follow_requests.receiver_id = $1 AND follow_requests.status = 'pending'
		)
		AND "user".user_id != $1
	`
	rows, err := DB.Query(context.Background(), query, userID)
	if err != nil {
		log.Printf("Database query error: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var username string
		if err = rows.Scan(&username); err != nil {
			log.Printf("database failed to scan explore user: %v\n", err)
			return nil, err
		}
		explore = append(explore, username)
	}

	if err = rows.Err(); err != nil {
		log.Printf("rows iteration error: %v\n", err)
		return nil, err
	}

	return explore, nil
}

func GetExploreGroup(groupID int)([]models.PostFeedProfile, error){
	var explore []models.PostFeedProfile
	query := `
		SELECT
		user_id,
		user_name,
		image,
		first_name,
		last_name
	FROM
		"user"
	INNER JOIN
		profile USING (user_id)
	WHERE
		user_id NOT IN (
			SELECT
				user_id
			FROM
				group_member
			WHERE
				group_id = $1
		)
	`
	rows, err := DB.Query(context.Background(), query, groupID)
	if err != nil {
		log.Printf("Database query error: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user models.PostFeedProfile
		if err = rows.Scan(&user.UserID, &user.UserName,&user.Avatar,&user.FirstName, &user.LastName); err != nil {
			log.Printf("database failed to scan explore user: %v\n", err)
			return nil, err
		}
		var invitationCount int
		query = `SELECT COUNT(*) FROM group_invitations WHERE receiver_id = $1 AND status = 'pending'`
		err = DB.QueryRow(context.Background(), query, user.UserID).Scan(&invitationCount)
		
		if err != nil {
			log.Printf("database failed to scan explore user: %v\n", err)
			return nil, err
		}
		if invitationCount == 0{
			user.UserID = ""
		explore = append(explore, user)
		}
	}

	if err = rows.Err(); err != nil {
		log.Printf("rows iteration error: %v\n", err)
		return nil, err
	}
	// log.Print(explore)
	return explore, nil
}
