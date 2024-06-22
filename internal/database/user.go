package database

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
	query := `INSERT INTO public.user (user_id, email, "password", provider)
	        VALUES ($1, $2, $3, $4)`
	// Execute SQL statement
	_, err := DB.Exec(context.Background(), query, user.UserID, user.Email, user.Password, user.Provider)
	if err != nil {
		log.Printf("Failed to insert user: %v\n", err)
		return err
	}
	return nil
}

func CreateUserProfile(userProfile models.UserProfile) error {
	// Prepare SQL statement
	query := `INSERT INTO public.profile (user_id, first_name, last_name, gender, date_of_birth, image, privacy, nick_name)
	        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`
	// Execute SQL statement
	_, err := DB.Exec(context.Background(), query, userProfile.UserID, userProfile.FirstName, userProfile.LastName,
		userProfile.Gender, userProfile.DateOfBirth, userProfile.Image, userProfile.Type, userProfile.NickName)
	if err != nil {
		log.Printf("Failed to insert user profile: %v\n", err)
		return err
	}
	return nil
}

func UpdateUserProfile(userProfile models.UserProfile) error {
	// Update user profile in database
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
	query := `SELECT user_id, email, provider FROM public.user WHERE user_id = $1`
	err := DB.QueryRow(context.Background(), query, userID).Scan(&user.UserID, &user.Email, &user.Provider)
	if err != nil {
		log.Printf("Failed to fetch user by ID: %v\n", err)
		return &user, err
	}
	return &user, nil
}
