package querys

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"social-network/internal/models"

	"github.com/gofrs/uuid"
)

func SignUpUser(user models.User, userProfile models.UserProfile) error {
	// Generate UUID if provider is "manual"
	if user.Provider == models.Provider.Manual {
		if user.Password == "" {
			return errors.New("password is required for manual registration")
		}
		// Assuming UserProfile also needs the same UserID
	}
	uuid, err := uuid.NewV4()
	if err != nil {
		log.Printf("Failed to generate UUID: %v\n", err)
		return err
	}
	user.UserID = uuid.String()
	userProfile.UserID = user.UserID
	// Add user to database
	if err := CreateUser(user); err != nil {
		return err
	}
	if err := CreateUserProfile(userProfile); err != nil {
		return err
	}
	return nil
}

func CreateUser(user models.User) error {
	// Prepare SQL statement
	query := `INSERT INTO public.user (user_id, user_name, email, "password", provider)
	        VALUES ($1, $2, $3, $4, $5)`
	// Execute SQL statement
	_, err := DB.Exec(context.Background(), query, user.UserID, user.UserName, user.Email, user.Password, user.Provider)
	if err != nil {
		log.Printf("Failed to insert user: %v\n", err)
		return err
	}
	return nil
}

func GetUserPostFeedProfile(userID string) (*models.PostFeedProfile, error) {
	// Fetch user profile from database
	var userProfile models.PostFeedProfile
	query := `
	SELECT
		public.profile.first_name,
		public.profile.last_name,
		public.profile.image,
		public.profile.privacy,
		public."user".user_name
	FROM
		public.profile
	INNER JOIN
		public."user"
	ON
		public.profile.user_id = public."user".user_id
	WHERE
		public.profile.user_id = $1`
	err := DB.QueryRow(context.Background(), query, userID).Scan(
		&userProfile.FirstName,
		&userProfile.LastName,
		&userProfile.Avatar,
		&userProfile.ProfilePrivacy,
		&userProfile.UserName,
	)
	if err != nil {
		log.Printf("Failed to fetch user profile: %v\n", err)
		return nil, err
	}

	return &userProfile, nil
}

func GetUserProfile(userID string) (*models.UserProfile, error) {
	// Fetch user profile from database
	var userProfile models.UserProfile
	query := `
	SELECT
		public.profile.first_name,
		public.profile.last_name,
		public.profile.gender,
		public.profile.date_of_birth,
		public.profile.image,
		public.profile.privacy,
		public.profile.nick_name,
		public.profile.about,
		public."user".user_name
	FROM
		public.profile
	INNER JOIN
		public."user"
	ON
		public.profile.user_id = public."user".user_id
	WHERE
		public.profile.user_id = $1`
	err := DB.QueryRow(context.Background(), query, userID).Scan(
		&userProfile.FirstName,
		&userProfile.LastName,
		&userProfile.Gender,
		&userProfile.DateOfBirth,
		&userProfile.Avatar,
		&userProfile.ProfilePrivacy,
		&userProfile.NickName,
		&userProfile.About,
		&userProfile.Username,
	)
	if err != nil {
		log.Printf("Failed to fetch user profile: %v\n", err)
		return nil, err
	}

	return &userProfile, nil
}

func GetUserProfileItem(userID string, item string) (interface{}, error) {
	// Fetch user profile from database
	var itemValue interface{}
	query := fmt.Sprintf(`SELECT %s FROM public.profile WHERE user_id = $1`, item)
	err := DB.QueryRow(context.Background(), query, userID).Scan(&itemValue)
	if err != nil {
		log.Printf("Failed to fetch user profile item: %v\n", err)
		return "", err
	}
	return itemValue, nil
}

func GetUserProfileByUserName(username string) (*models.UserProfile, error) {
	// Fetch user profile by username from database
	var userProfile models.UserProfile
	query := `
	SELECT
		public.profile.first_name,
		public.profile.last_name,
		public.profile.gender,
		public.profile.date_of_birth,
		public.profile.image,
		public.profile.privacy,
		public.profile.nick_name,
		public.profile.about,
		public.profile.user_id,
		public."user".user_name
	FROM
		public.profile
	INNER JOIN
		public."user"
	ON
		public.profile.user_id = public."user".user_id
	WHERE
		public."user".user_name = $1`
	err := DB.QueryRow(context.Background(), query, username).Scan(
		&userProfile.FirstName,
		&userProfile.LastName,
		&userProfile.Gender,
		&userProfile.DateOfBirth,
		&userProfile.Avatar,
		&userProfile.ProfilePrivacy,
		&userProfile.NickName,
		&userProfile.About,
		&userProfile.UserID,
		&userProfile.Username,
	)
	if err != nil {
		log.Printf("Failed to fetch user profile by username: %v\n", err)
		return nil, err
	}

	return &userProfile, nil
}

func IsPrivateUser(userID string) (bool, error) {
	var privacy string
	query := `SELECT privacy FROM public.profile WHERE user_id = $1`
	err := DB.QueryRow(context.Background(), query, userID).Scan(&privacy)
	if err != nil {
		log.Printf("Failed to fetch user profile: %v\n", err)
		return false, err
	}
	return privacy == models.ProfilePrivacy.Private, nil
}

func CreateUserProfile(userProfile models.UserProfile) error {
	// Prepare SQL statement
	query := `INSERT INTO public.profile (user_id, first_name, last_name, gender, date_of_birth, image, privacy, nick_name, about)
	        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	// Execute SQL statement
	_, err := DB.Exec(context.Background(), query,
		userProfile.UserID,
		userProfile.FirstName,
		userProfile.LastName,
		userProfile.Gender,
		userProfile.DateOfBirth,
		userProfile.Avatar,
		userProfile.ProfilePrivacy,
		userProfile.NickName,
		userProfile.About,
	)
	if err != nil {
		log.Printf("Failed to insert user profile: %v\n", err)
		return err
	}
	return nil
}

func UpdateUserProfile(userProfile models.UserProfile) error {
	// Prepare SQL statement
	query := `UPDATE public.profile SET
		first_name = $1,
		last_name = $2,
		gender = $3,
		date_of_birth = $4,
		image = $5,
		privacy = $6,
		nick_name = $7,
		about = $8
	WHERE user_id = $9`
	// Execute SQL statement
	_, err := DB.Exec(context.Background(), query,
		userProfile.FirstName,
		userProfile.LastName,
		userProfile.Gender,
		userProfile.DateOfBirth,
		userProfile.Avatar,
		userProfile.ProfilePrivacy,
		userProfile.NickName,
		userProfile.About,
		userProfile.UserID,
	)
	if err != nil {
		log.Printf("Failed to update user profile: %v\n", err)
		return err
	}
	return nil
}

func GetManualUser(email string) (*models.User, error) {
	// Fetch user password hash from database
	var user models.User
	query := `SELECT user_id, email, password, provider FROM public.user WHERE email = $1 AND provider = 'manual'`
	err := DB.QueryRow(context.Background(), query, email).Scan(&user.UserID, &user.Email, &user.Password, &user.Provider)
	if err != nil {
		log.Printf("Failed to fetch manual user: %v\n", err)
		return &user, err
	}
	return &user, nil
}

func AddUserSession(userID string, sessionID string) error {
	// Add user session to database
	query := `INSERT INTO public.session (user_id, session_uuid) VALUES ($1, $2)`
	_, err := DB.Exec(context.Background(), query, userID, sessionID)
	if err != nil {
		log.Printf("Failed to insert user session: %v\n", err)
		return err
	}
	return nil
}

	// delete the previous session of the user if exists
	func DeleteUserSessions(userID string) error {
		query := `DELETE FROM public.session WHERE user_id = $1`
		_, err := DB.Exec(context.Background(), query, userID)
		if err != nil {
			log.Printf("Failed to delete user session: %v\n", err)
			return err
		}
		return nil
	}

func DeleteUserSession(sessionID string) error {
	// Delete user session from database
	query := `DELETE FROM public.session WHERE session_uuid = $1`
	_, err := DB.Exec(context.Background(), query, sessionID)
	if err != nil {
		log.Printf("Failed to delete user session: %v\n", err)
		return err
	}
	return nil
}

func ValidateSessionToken(sessionID string) (string, error) {
	// Validate session token and return user ID
	var userID string
	query := `SELECT user_id FROM public.session WHERE session_uuid = $1`
	err := DB.QueryRow(context.Background(), query, sessionID).Scan(&userID)
	if err != nil {
		log.Printf("Failed to validate session token: %v\n", err)
		return "", err
	}
	return userID, nil
}

func GetUserByID(userID string) (*models.User, error) {
	// Fetch user by user ID
	var user models.User
	query := `SELECT user_id, user_name, email, provider FROM public.user WHERE user_id = $1`
	err := DB.QueryRow(context.Background(), query, userID).Scan(&user.UserID, &user.UserName, &user.Email, &user.Provider)
	if err != nil {
		log.Printf("Failed to fetch user by ID: %v\n", err)
		return &user, err
	}
	return &user, nil
}

func GetUserGroups(userID string) (map[int]bool, error) {
	Groups := make(map[int]bool, 0)
	query := `SELECT group_id FROM group_member WHERE user_id = $1`
	rows, err := DB.Query(context.Background(), query, userID)
	if err != nil {
		log.Printf("database failed to scan followed user: %v\n", err)
		return nil, err
	}
	for rows.Next() {
		var groupID int
		if err = rows.Scan(&groupID); err != nil {
			log.Printf("database failed to scan followed user: %v\n", err)
			return nil, err
		}
		Groups[groupID] = true
	}
	return Groups, nil
}

// get a userID by uerName
func GetUserIDByUserName(userName string) (string, error) {
	var userID string
	query := `SELECT user_id FROM public.user WHERE user_name = $1`
	err := DB.QueryRow(context.Background(), query, userName).Scan(&userID)
	if err != nil {
		log.Printf("Failed to fetch user by username: %v\n", err)
		return "", err
	}
	return userID, nil
}

// IsFollowing checks if a user is following another user
func IsFollowing(userID string, followedID string) bool {
	query := `SELECT followed_id FROM follower WHERE follower_id = $1 AND followed_id = $2`
	err := DB.QueryRow(context.Background(), query, userID, followedID).Scan(&followedID)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return false
		}
		log.Printf("database: Failed to check if user is following: %v", err)
		return false
	}
	return true
}

func GetRequestStatus(userID string, followedID string) (string, error) {
	var status string
	query := `SELECT status FROM follow_requests WHERE sender_id = $1 AND receiver_id = $2 AND status = 'pending'`
	err := DB.QueryRow(context.Background(), query, userID, followedID).Scan(&status)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return "not_requested", err
		}
		log.Printf("database: Failed to check if user is following: %v", err)
		return "", err
	}
	return status, nil
}

func GetFollowStatus(userID string, followedID string) string {
	if IsFollowing(userID, followedID) || userID == followedID {
		return "following"
	}
	status, err := GetRequestStatus(userID, followedID)
	if err != nil {
		return "not_following"
	}
	if status == "pending" {
		return "pending"
	}
	return "not_following"
}

// FollowUser adds a follow relationship between two users
func FollowUser(request models.Request) error {
	// Check if user is already following
	if IsFollowing(request.Sender, request.Receiver) {
		return errors.New("user is already following")
	}
	// Insert follow relationship into database
	query := `INSERT INTO follower (follower_id, followed_id) VALUES ($1, $2)`
	_, err := DB.Exec(context.Background(), query, request.Sender, request.Receiver)
	if err != nil {
		log.Printf("database: Failed to insert follow relationship: %v", err)
		return err
	}
	return nil
}

// UnFollowUser removes a follow relationship between two users
func UnFollowUser(request models.Request) error {
	// Check if user is following
	if !IsFollowing(request.Sender, request.Receiver) {
		return errors.New("user is not following")
	}
	// Remove follow relationship from database
	query := `DELETE FROM follower WHERE follower_id = $1 AND followed_id = $2`
	_, err := DB.Exec(context.Background(), query, request.Sender, request.Receiver)
	if err != nil {
		log.Printf("database: Failed to delete follow relationship: %v", err)
		return err
	}
	return nil
}

// get userName by ID
func GetUserNameByID(userID string) (string, error) {
	var username string
	query := `SELECT user_name FROM public.user WHERE user_id = $1`
	err := DB.QueryRow(context.Background(), query, userID).Scan(&username)
	if err != nil {
		log.Printf("Failed to fetch user by ID: %v\n", err)
		return "", err
	}
	return username, nil
}

// Function to check if a username exists in the database
func usernameExists(username string) (bool, error) {
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM public.user WHERE user_name=$1)`
	err := DB.QueryRow(context.Background(), query, username).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

// Function to generate a unique username
func GenerateUniqueUsername(firstName, lastName string) (string, error) {
	// Create a base username
	baseUsername := strings.ToLower(fmt.Sprintf("%s.%s", firstName, lastName))

	// Ensure the base username is within the 10 character limit
	if len(baseUsername) > 10 {
		baseUsername = baseUsername[:10]
	}

	username := baseUsername
	counter := 1

	// Check if the username exists and generate new variations if necessary
	for {
		exists, err := usernameExists(username)
		if err != nil {
			return "", err
		}
		if !exists {
			break
		}
		username = fmt.Sprintf("%s%d", baseUsername, counter)
		counter++
	}

	return username, nil
}

func GetUserIDByProvider(user models.User, UserProfile models.UserProfile) (string, error) {
	query := `SELECT user_id FROM public.user WHERE email = $1 and provider = $2`
	// Execute the query
	var userID string
	DB.QueryRow(context.Background(), query, user.Email, user.Provider).Scan(&userID)
	if userID == "" {
		// Add the user to the database
		if err := SignUpUser(user, UserProfile); err != nil {
			log.Printf("Error signing up user: %v", err)
		}
		DB.QueryRow(context.Background(), query, user.Email, user.Provider).Scan(&userID)
	}
	sessionID, err := CreateLoginSession(userID)
	return sessionID, err
}

func CreateLoginSession(userID string) (string, error) {
	sessionUUID, err := uuid.NewV4()
	if err != nil {
		log.Printf("Error creating session UUID: %v", err)
		return "", err
	}
	if err := AddUserSession(userID, sessionUUID.String()); err != nil {
		log.Printf("Error adding session: %v", err)
		return "", err
	}
	return sessionUUID.String(), nil
}
