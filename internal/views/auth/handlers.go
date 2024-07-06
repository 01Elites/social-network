package auth

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	database "social-network/internal/database/querys"
	"social-network/internal/helpers"
	"social-network/internal/models"
	"social-network/internal/views/session"

	"math/rand"

	"github.com/gofrs/uuid"
)

type SignUpRequst struct {
	UserName       string    `json:"user_name"`
	Email          string    `json:"email"`
	Password       string    `json:"password"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	Gender         string    `json:"gender"`
	NickName       string    `json:"nick_name"`
	ProfilePrivacy string    `json:"profile_privacy"`
	Image          string    `json:"image"`
	About          string    `json:"about"`
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	var data SignUpRequst

	// Create a JSON decoder for the request body
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	// Decode the JSON directly from the request body
	err := decoder.Decode(&data)
	if err != nil {
		log.Printf("Error decoding sign up request: %v", err)
		helpers.HTTPError(w, "Something Went Wrong!!", http.StatusBadRequest)
		return
	}

	if err := ValidatePassword(data.Password); err != nil {
		log.Printf("Password validation error: %v", err)
		helpers.HTTPError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if data.ProfilePrivacy == "" {
		data.ProfilePrivacy = models.ProfilePrivacy.Public
	}

	if data.Image != "" { // Save the image if it exists
		data.Image, err = helpers.SaveBase64Image(data.Image)
		if err != nil {
			fmt.Println("Error with Profile Image:\n", err)
		}
	}

	userProfile := models.UserProfile{
		NickName:       data.NickName,
		FirstName:      data.FirstName,
		LastName:       data.LastName,
		Gender:         data.Gender,
		DateOfBirth:    data.DateOfBirth,
		ProfilePrivacy: data.ProfilePrivacy,
		Image:          data.Image,
		About:          data.About,
	}

	if err := helpers.ValidateUserProfileData(&userProfile); err != nil {
		log.Printf("Validation error: %v", err)
		helpers.HTTPError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Hash the password before adding it
	hash, err := HashPassword(data.Password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		helpers.HTTPError(w, "Something Went Wrong!!", http.StatusInternalServerError)
		return
	}
	data.Password = hash
	data.Email = strings.ToLower(data.Email) // Convert email to lowercase for consistency

	// -------- GENERATE A RANDOM USERNAME ---------
	rand.Seed(time.Now().UnixNano())
	randomDigits := fmt.Sprintf("%04d", rand.Intn(10000))

	user := models.User{
		UserName: data.FirstName[:1] + data.LastName + randomDigits,
		Email:    data.Email,
		Password: data.Password,
		Provider: models.Provider.Manual,
	}

	if err := helpers.ValidateUserFields(&user); err != nil {
		log.Printf("Validation error: %v", err)
		helpers.HTTPError(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := database.SignUpUser(user, userProfile); err != nil {
		if strings.Contains(err.Error(), "SQLSTATE 23505") {
			if strings.Contains(err.Error(), "user_user_name_key") {
				log.Printf("Username already exists: %v", err)
				helpers.HTTPError(w, "Username already exists", http.StatusConflict)
				return
			} else if strings.Contains(err.Error(), "unique_email_provider") {
				log.Printf("Email already exists: %v", err)
				helpers.HTTPError(w, "Email already exists", http.StatusConflict)
				return
			}
			log.Printf("User already exists: %v", err)
			helpers.HTTPError(w, "User already exists", http.StatusConflict)
			return
		} else if strings.Contains(err.Error(), "SQLSTATE 22P02") {
			log.Printf("Invalid data format: %v", err)
			helpers.HTTPError(w, "Invalid data format", http.StatusBadRequest)
			return
		}
		log.Printf("Error signing up user: %v", err)
		helpers.HTTPError(w, "Internal Server error", http.StatusInternalServerError)
		return
	}

	// AddClient(data.UserName)
	// SendUsersList()
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Signup successful")
}

type SignInRequst struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var data SignInRequst

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	if err := decoder.Decode(&data); err != nil {
		log.Printf("Error decoding sign up request: %v", err)
		helpers.HTTPError(w, "Something Went Wrong!!", http.StatusBadRequest)
		return
	}

	data.Email = strings.ToLower(data.Email) // Convert email to lowercase for consistency

	user, err := database.GetManualUser(data.Email)
	if err != nil {
		log.Printf("Error getting manual user: %v", err)
		helpers.HTTPError(w, "User not found", http.StatusBadRequest)
		return
	}
	hashMatched := CheckPasswordHash(data.Password, user.Password)
	if !hashMatched {
		log.Printf("Password does not match")
		helpers.HTTPError(w, "username or password is incorrect", http.StatusUnauthorized)
		return
	}
	// Create a seesion for this user
	sessionUUID, err := uuid.NewV4()
	if err != nil {
		log.Printf("Error creating session UUID: %v", err)
		helpers.HTTPError(w, "Something Went Wrong!!", http.StatusInternalServerError)
		return
	}

	// Add session to database
	if err := database.AddUserSession(user.UserID, sessionUUID.String()); err != nil {
		log.Printf("Error adding session: %v", err)
		helpers.HTTPError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Set a cookie with a session token that can be used to authenticate access without logging in
	// session.SetSessionCookie(w, sessionUUID.String())

	session.SetAutherizationHeader(w, sessionUUID.String())
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Signin successful")
}

func LogOut(w http.ResponseWriter, r *http.Request) { // Get the session token from the cookie
	token, err := session.ExtractToken(r)
	if err != nil {
		log.Printf("Error extracting token: %v", err)
		helpers.HTTPError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Delete the session from the database
	if err := database.DeleteUserSession(token); err != nil {
		log.Printf("Error deleting session: %v", err)
		helpers.HTTPError(w, "Internal Server error", http.StatusInternalServerError)
		return
	}
	// Expire the cookie
	session.ClearAutherizationHeader(w)
	// AddClient(data.UserName)
	io.WriteString(w, "LogOut success")
}
