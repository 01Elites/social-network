package profile

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"social-network/internal/helpers"
	"social-network/internal/models"
	"social-network/internal/views/middleware"

	database "social-network/internal/database/querys"
)

type profileData struct {
	UserName       string    `json:"user_name"`
	Email          string    `json:"email"`
	NickName       string    `json:"nick_name"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Gender         string    `json:"gender"`
	DateOfBirth    time.Time `json:"date_of_birth"`
	AvatarURL      string    `json:"avatar_url"`
	About          string    `json:"about"`
	ProfilePrivacy string    `json:"profile_privacy"`
	Follow_status  string    `json:"follow_status"`
}

func getProfile(w http.ResponseWriter, r *http.Request) {
	// Retrieve the userID from context using the same key defined globally
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		log.Printf("User ID not found in get profile handler\n")
		helpers.HTTPError(w, "User ID not found", http.StatusInternalServerError)
		return
	}

	var profile profileData
	user, err := database.GetUserByID(userID)
	if err != nil {
		log.Printf("Failed to get user details: %v\n", err)
		helpers.HTTPError(w, "Failed to get user details", http.StatusInternalServerError)
		return
	}
	prof, err := database.GetUserProfile(userID)
	if err != nil {
		log.Printf("Error getting user profile: %v\n", err)
		helpers.HTTPError(w, "Failed to get user profile", http.StatusInternalServerError)
		return
	}
	profile = profileData{
		UserName:       user.UserName,
		Email:          user.Email,
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
		log.Printf("Failed to encode profile data: %v\n", err)
		helpers.HTTPError(w, "Something Went Wrong!!", http.StatusInternalServerError)
		return
	}
}

func patchProfile(w http.ResponseWriter, r *http.Request) {
	// Retrieve the userID from context using the same key defined globally
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		log.Printf("User ID not found in patch profile handler\n")
		helpers.HTTPError(w, "Something Went Wrong!!", http.StatusInternalServerError)
		return
	}
	var update profileData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&update)
	if err != nil {
		log.Printf("Error decoding request body: %v\n", err)
		helpers.HTTPError(w, "Something Went Wrong!!", http.StatusBadRequest)
		return
	}

	if update.AvatarURL != "" { // If the user has no image, use the default image
		update.AvatarURL, err = helpers.SaveBase64Image(update.AvatarURL)
		if err != nil {
			fmt.Println("Error with Image:", err)
		}
	}

	updatedProfile := models.UserProfile{
		UserID:         userID,
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
		log.Printf("Validation error: %v\n", err)
		helpers.HTTPError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update the user profile
	if err := database.UpdateUserProfile(updatedProfile); err != nil {
		log.Printf("Failed to update user profile: %v\n", err)
		helpers.HTTPError(w, "Something Went Wrong!!", http.StatusInternalServerError)
		return
	}

	io.WriteString(w, "Profile updated successfully")
}

// getProfileByID retrieves the profile of a user by their ID
func getProfileByUserName(w http.ResponseWriter, r *http.Request) {
	// Retrieve the userID from context using the same key defined globally
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		log.Printf("User ID not found")
		helpers.HTTPError(w, "User ID not found", http.StatusInternalServerError)
		return
	}

	userName := r.PathValue("user_name")
	UserPageID, err := database.GetUserIDByUserName(userName)
	if err != nil {
		log.Printf("Failed to get userID: %v\n", err)
		helpers.HTTPError(w, "Failed to get userID", http.StatusInternalServerError)
		return
	}

	var profile profileData
	user, err := database.GetUserByID(UserPageID)
	if err != nil {
		log.Printf("Failed to get user: %v\n", err)
		helpers.HTTPError(w, "Failed to get user", http.StatusInternalServerError)
		return
	}
	prof, err := database.GetUserProfile(UserPageID)
	if err != nil {
		log.Printf("Error getting user profile: %v\n", err)
		helpers.HTTPError(w, "Failed to get user profile", http.StatusInternalServerError)
		return
	}

	profile = profileData{
		UserName: user.UserName,
		// Email:     user.Email,
		NickName:  prof.NickName,
		FirstName: prof.FirstName,
		LastName:  prof.LastName,
		// Gender:         prof.Gender,
		// DateOfBirth:    prof.DateOfBirth,
		AvatarURL:      prof.Image,
		About:          prof.About,
		ProfilePrivacy: prof.ProfilePrivacy,
		Follow_status:  database.GetFollowStatus(userID, UserPageID),
	}
	if err := json.NewEncoder(w).Encode(profile); err != nil {
		log.Printf("Failed to encode profile data: %v\n", err)
		helpers.HTTPError(w, "Something Went Wrong!!", http.StatusInternalServerError)
		return
	}
}

func getProfilePosts(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		log.Printf("User ID not found in get profile posts handler\n")
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}

	userName := r.PathValue("user_name")
	UserPageID, err := database.GetUserIDByUserName(userName)
	if err != nil {
		log.Printf("Failed to get userID: %v\n", err)
		helpers.HTTPError(w, "Failed to get userID", http.StatusInternalServerError)
		return
	}

	posts, err := database.GetUserPosts(userID, UserPageID, database.IsFollowing(userID, UserPageID))
	if err != nil {
		log.Printf("Failed to get posts: %v\n", err)
		helpers.HTTPError(w, "Failed to get posts", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}
