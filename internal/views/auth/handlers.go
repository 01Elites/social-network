package auth

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"

	"social-network/internal/database"
	"social-network/internal/models"

	"github.com/gofrs/uuid"
)

type SignUpRequst struct {
	Email          string `json:"email"`
	Password       string `json:"password"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	DateOfBirth    string `json:"date_of_birth"`
	Gender         string `json:"gender"`
	NickName       string `json:"nick_name"`
	ProfilePrivacy string `json:"profile_privacy"`
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	var data SignUpRequst

	// Create a JSON decoder for the request body
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	// Decode the JSON directly from the request body
	if err := decoder.Decode(&data); err != nil {
		log.Printf("Error decoding sign up request: %v", err)
		SendError(w, "Failed to decode JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := ValidateSignUpData(&data); err != nil {
		log.Printf("Validation error: %v", err)
		SendError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Hash the password before adding it
	hash, err := HashPassword(data.Password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		SendError(w, "Internal error", http.StatusInternalServerError)
		return
	}
	data.Password = hash
	user := models.User{
		Email:    data.Email,
		Password: data.Password,
		Provider: models.Provider.Manual,
	}
	userProfile := models.UserProfile{
		NickName:    data.NickName,
		FirstName:   data.FirstName,
		LastName:    data.LastName,
		Gender:      data.Gender,
		DateOfBirth: data.DateOfBirth,
		// Image:       "",
		Type: data.ProfilePrivacy,
	}
	if err := database.SignUpUser(user, userProfile); err != nil {
		if strings.Contains(err.Error(), "SQLSTATE 23505") {
			log.Printf("User already exists: %v", err)
			SendError(w, "User already exists", http.StatusConflict)
			return
		} else if strings.Contains(err.Error(), "SQLSTATE 22P02") {
			log.Printf("Invalid data format: %v", err)
			SendError(w, "Invalid data format", http.StatusBadRequest)
			return
		}
		log.Printf("Error signing up user: %v", err)
		SendError(w, "Internal server error", http.StatusInternalServerError)
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
		SendError(w, "Failed to decode JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	user, err := database.GetManualUser(data.Email)
	if err != nil {
		log.Printf("Error getting user hash: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	hashMatched := CheckPasswordHash(data.Password, user.Password)
	if !hashMatched {
		log.Printf("Password does not match")
		SendError(w, "username or password is incorrect", http.StatusUnauthorized)
		return
	}
	// Create a seesion for this user
	sessionUUID, err := uuid.NewV4()
	if err != nil {
		SendError(w, "Failed to create session", http.StatusInternalServerError)
		return
	}

	// Add session to database
	if err := database.AddUserSession(user.UserID, sessionUUID.String()); err != nil {
		SendError(w, "Failed to add session", http.StatusInternalServerError)
		return
	}

	// Set a cookie with a session token that can be used to authenticate access without logging in
	SetSessionCookie(w, sessionUUID.String())

	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Signin successful")
}

func LogOut(w http.ResponseWriter, r *http.Request) { // Get the session token from the cookie
	cookie, err := r.Cookie("SN_SESSION")
	if err != nil {
		// No session token, user is not logged in
		return
	}
	sessionToken := cookie.Value

	// Delete the session from the database
	if err := database.DeleteUserSession(sessionToken); err != nil {
		SendError(w, "Failed to delete session", http.StatusInternalServerError)
		return
	}
	// Expire the cookie
	clearSessionCookie(w)
	// AddClient(data.UserName)
	io.WriteString(w, "LogOut success")
}

type ErrorResponse struct {
	Reason string `json:"reason"`
}

func SendError(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	errResponse := ErrorResponse{Reason: message}
	json.NewEncoder(w).Encode(errResponse)
}
