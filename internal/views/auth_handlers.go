package views

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"social-network/internal/database"
	"social-network/internal/models"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
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
		http.Error(w, "Failed to decode JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	if err := ValidateSignUpData(&data); err != nil {
		log.Printf("Validation error: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Hash the password before adding it
	hash, err := HashPassword(data.Password)
	if err != nil {
		log.Printf("Error hashing password: %v", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
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
			http.Error(w, "User already exists", http.StatusConflict)
			return
		} else if strings.Contains(err.Error(), "SQLSTATE 22P02") {
			log.Printf("Invalid data format: %v", err)
			http.Error(w, "Invalid data format", http.StatusBadRequest)
			return
		}
		log.Printf("Error signing up user: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
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
		http.Error(w, "Failed to decode JSON: "+err.Error(), http.StatusBadRequest)
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
		http.Error(w, "username or password is incorrect", http.StatusUnauthorized)
		return
	}
	// Create a seesion for this user
	sessionUUID, err := uuid.NewV4()
	if err != nil {
		http.Error(w, "Failed to create session", http.StatusInternalServerError)
		return
	}

	// Add session to database
	if err := database.AddUserSession(user.UserID, sessionUUID.String()); err != nil {
		http.Error(w, "Failed to add session", http.StatusInternalServerError)
		return
	}

	// Set a cookie with a session token that can be used to authenticate access without logging in
	setSessionCookie(w, sessionUUID.String())

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
		http.Error(w, "Failed to delete session", http.StatusInternalServerError)
		return
	}
	// Expire the cookie
	clearSessionCookie(w)
	// AddClient(data.UserName)
	io.WriteString(w, "LogOut success")
}

func validateSession(w http.ResponseWriter, r *http.Request) (*models.User, error) {
	// Extract the session token from the cookie
	cookie, err := r.Cookie("SN_SESSION")
	if err != nil {
		return nil, err // No cookie means no session
	}
	// Validate the session token in the database
	userID, err := database.ValidateSessionToken(cookie.Value)
	if err != nil {
		return nil, err // Invalid or expired session token
	}

	// Retrieve user details based on userID
	user, err := database.GetUserByID(userID)
	if err != nil {
		return nil, err
	}

	setSessionCookie(w, cookie.Value)
	return user, nil
}

/********************* Session Cookie ************************/

func clearSessionCookie(w http.ResponseWriter) {
	// Clear the session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "SN_SESSION",
		Value:    "",
		Expires:  time.Now().Add(-10 * time.Second), // to enhance the compatibility and ensure all browsers handle the cookie clearing
		Path:     "/",
		MaxAge:   -1,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
	})
}

// updateSessionCookie updates the session cookie expiration time.
func setSessionCookie(w http.ResponseWriter, sessionToken string) {
	expiration := time.Now().AddDate(1, 0, 0)
	updatedCookie := http.Cookie{
		Name:     "SN_SESSION",
		Value:    sessionToken,
		Expires:  expiration,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
		HttpOnly: true,
	}
	http.SetCookie(w, &updatedCookie)
}

/********************* Validation ************************/

// ValidateSignUpData validates the sign-up data and returns an error if any validation fails.
func ValidateSignUpData(data *SignUpRequst) error {
	// Trim spaces and check for empty required fields
	data.Email = strings.TrimSpace(data.Email)
	data.NickName = strings.TrimSpace(data.NickName)
	data.FirstName = strings.TrimSpace(data.FirstName)
	data.LastName = strings.TrimSpace(data.LastName)

	if data.Email == "" || data.NickName == "" || data.FirstName == "" || data.LastName == "" {
		return errors.New("all fields must be filled out")
	}

	// Check the length of the username
	if len(data.NickName) > 20 {
		return errors.New("nickname should be up to 20 characters long")
	}

	// Check the length of the first name and last name
	if len(data.FirstName) > 20 || len(data.LastName) > 20 {
		return errors.New("first name and last name should be up to 20 characters long")
	}

	// Check the length of the email
	if len(data.Email) > 30 {
		return errors.New("email should be up to 30 characters long")
	}

	// Validate password
	if checkPassWS(data.Password) {
		return errors.New("password cannot contain whitespaces")
	}

	if len(data.Password) < 6 || len(data.Password) > 20 {
		return errors.New("password should be between 6 and 20 characters long")
	}

	// Validate Date of Birth to check if user is at least 18 years old
	if err := validateAge(data.DateOfBirth); err != nil {
		return err
	}

	// Validate privacy
	switch data.ProfilePrivacy {
	case models.ProfilePrivacy.Private, models.ProfilePrivacy.Public:
		// valid
	default:
		return fmt.Errorf("invalid profile privacy: %v", data.ProfilePrivacy)
	}

	return nil
}

// checkPassWS checks if the given password string contains any whitespace characters.
// It returns true if there are whitespace characters, and false otherwise.
func checkPassWS(pass string) bool {
	return strings.ContainsAny(pass, " \t\n\r\v\f")
}

// validateAge checks if the user is at least 18 years old based on their date of birth.
func validateAge(dob string) error {
	// Parse the date of birth string into a time.Time object
	dateOfBirth, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return fmt.Errorf("invalid date format: %v", err)
	}

	// Get the current date
	now := time.Now()

	// Calculate the user's age
	age := now.Year() - dateOfBirth.Year()
	if now.YearDay() < dateOfBirth.YearDay() {
		age--
	}

	// Check if the user is at least 18 years old
	if age < 18 {
		return errors.New("you must be at least 18 years old to sign up")
	}

	return nil
}

/********************* hash password ************************/

// HashPassword hashes the given password using bcrypt and returns the hashed password.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash compares the given password with the given hash and returns true if they match, false otherwise.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
