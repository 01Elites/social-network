package querys

import (
	"context"
	"errors"
	"log"

	"social-network/internal/models"

	"github.com/gofrs/uuid"
)

func SignUpUser(user models.User, userProfile models.UserProfile) error {
	// Generate UUID if provider is "manual"
	if user.Provider == models.Provider.Manual {
		if user.Password == "" {
			return errors.New("password is required for manual registration")
		}
		uuid, err := uuid.NewV4()
		if err != nil {
			log.Printf("Failed to generate UUID: %v\n", err)
			return err
		}
		user.UserID = uuid.String()
		userProfile.UserID = user.UserID // Assuming UserProfile also needs the same UserID
	}
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

func GetUserEmailUserName(userID string) (*models.User, error) {
	query := `SELECT email, user_name FROM public.user WHERE user_id = $1`
	user := &models.User{}
	err := DB.QueryRow(context.Background(), query, userID).Scan(&user.Email, &user.UserName)
	if err != nil {
		log.Printf("Failed to fetch user email: %v\n", err)
		return nil, err
	}
	return user, nil
}

func GetUserProfile(userID string) (*models.UserProfile, error) {
	// Fetch user profile from database
	var userProfile models.UserProfile
	query := `SELECT first_name, last_name, gender, date_of_birth, image, privacy, nick_name, about FROM public.profile WHERE user_id = $1`
	err := DB.QueryRow(context.Background(), query, userID).Scan(
		&userProfile.FirstName,
		&userProfile.LastName,
		&userProfile.Gender,
		&userProfile.DateOfBirth,
		&userProfile.Image,
		&userProfile.ProfilePrivacy,
		&userProfile.NickName,
		&userProfile.About,
	)
	if err != nil {
		log.Printf("Failed to fetch user profile: %v\n", err)
		return nil, err
	}

	return &userProfile, nil
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
		userProfile.Image,
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
		userProfile.Image,
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
		log.Printf("Failed to fetch user by ID: %v\n", err)
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