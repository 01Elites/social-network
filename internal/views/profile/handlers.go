package profile

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"social-network/internal/database"
	"social-network/internal/helpers"
	"social-network/internal/models"
	"social-network/internal/views/middleware"
)

// ProfileHandler checks the HTTP method and delegates to the correct function
func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getProfile(w, r)
	case "PATCH":
		patchProfile(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

type profileData struct {
	Email          string    `json:"email"`
	NickName       string    `json:"nick_name"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Gender         string    `json:"gender"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	AvatarURL      string    `json:"avatar_url"`
	About          string    `json:"about"`
	ProfilePrivacy string    `json:"profile_privacy"`
}

func getProfile(w http.ResponseWriter, r *http.Request) {
	// Retrieve the userID from context using the same key defined globally
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		helpers.HTTPError(w, "User ID not found", http.StatusInternalServerError)
		return
	}

	var profile profileData
	email, err := database.GetUserEmail(userID)
	if err != nil {
		helpers.HTTPError(w, "Failed to get user email", http.StatusInternalServerError)
		return
	}
	prof, err := database.GetUserProfile(userID)
	if err != nil {
		helpers.HTTPError(w, "Failed to get user profile", http.StatusInternalServerError)
		return
	}
	profile = profileData{
		Email:          email,
		NickName:       prof.NickName,
		FirstName:      prof.FirstName,
		LastName:       prof.LastName,
		Gender:         prof.Gender,
		DateOfBirth:    prof.DateOfBirth,
		AvatarURL:      prof.Image,
		About:          prof.About,
		ProfilePrivacy: prof.ProfilePrivacy,
	}
	if err := json.NewEncoder(w).Encode(profile); err != nil {
		log.Printf("Failed to encode profile data: %v", err)
		helpers.HTTPError(w, "Something Went Wrong!!", http.StatusInternalServerError)
		return
	}
}

func patchProfile(w http.ResponseWriter, r *http.Request) {
	// Retrieve the userID from context using the same key defined globally
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		log.Printf("User ID not found")
		helpers.HTTPError(w, "Something Went Wrong!!", http.StatusInternalServerError)
		return
	}
	var update profileData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&update)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		helpers.HTTPError(w, "Something Went Wrong!!", http.StatusBadRequest)
		return
	}
	updatedProfile := models.UserProfile{
		NickName:       update.NickName,
		FirstName:      update.FirstName,
		LastName:       update.LastName,
		Gender:         update.Gender,
		DateOfBirth:    update.DateOfBirth,
		Image:          update.AvatarURL,
		About:          update.About,
		ProfilePrivacy: update.ProfilePrivacy,
	}
	// Validate the profile data
	if err := helpers.ValidateUserProfileData(&updatedProfile); err != nil {
		log.Printf("Validation error: %v", err)
		helpers.HTTPError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update the user profile
	if err := database.UpdateUserProfile(userID, updatedProfile); err != nil {
		log.Printf("Failed to update user profile: %v", err)
		helpers.HTTPError(w, "Something Went Wrong!!", http.StatusInternalServerError)
		return
	}

	io.WriteString(w, "Profile updated successfully")
}
