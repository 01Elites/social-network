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
	Email          string    `json:"email,omitempty"`
	NickName       string    `json:"nick_name"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Gender         string    `json:"gender,omitempty"`
	DateOfBirth    time.Time `json:"date_of_birth,omitempty"`
	Avatar         string    `json:"avatar"`
	About          string    `json:"about"`
	ProfilePrivacy string    `json:"profile_privacy"`
	Follow_status  string    `json:"follow_status,omitempty"`
	FollowingCount int       `json:"following_count,omitempty"`
	FollowerCount  int       `json:"follower_count,omitempty"`
}

/*
getProfile handles the HTTP GET request for retrieving a user's profile.
It reads the user ID from the request context, fetches the user's details
and profile data from the database, and returns the profile in JSON format.

Endpoint: GET /api/profile

Response:

	{
		"user_name": string,
		"email": string,
		"nick_name": string,
		"first_name": string,
		"last_name": string,
		"gender": "male" | "female",
		"date_of_birth": 0, // unix timestamp
		"avatar": string,
		"about": string,
		"profile_privacy": "public" | "private",
	}

If any error occurs during the process, it returns the corresponding HTTP error status code.
*/
func getMyProfile(w http.ResponseWriter, r *http.Request) {
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
		Avatar:         prof.Avatar,
		About:          prof.About,
		ProfilePrivacy: prof.ProfilePrivacy,
	}
	if err := json.NewEncoder(w).Encode(profile); err != nil {
		log.Printf("Failed to encode profile data: %v\n", err)
		helpers.HTTPError(w, "Something Went Wrong!!", http.StatusInternalServerError)
		return
	}
}

/*
patchProfile handles the HTTP PATCH request for updating a user's profile.
It reads the user ID from the request context, unmarshals the JSON request body,
validates the profile data, updates the user's profile in the database, and returns an appropriate response.

Endpoint: PATCH /api/profile

Request Body:

	{
		"nick_name": string, // optional
		"first_name": string,
		"last_name": string,
		"gender": "male" | "female",
		"date_of_birth": string, // ISO 8601 date string
		"avatar": string, // Base64 encoded image, optional
		"about": string, // optional
		"profile_privacy": "public" | "private"
	}

If any error occurs during the process, it returns the corresponding HTTP error status code.
*/
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

	if update.Avatar != "" { // If the user has no image, use the default image
		update.Avatar, err = helpers.SaveBase64Image(update.Avatar)
		if err != nil {
			fmt.Println("Error with Image:", err)
		}
	} else {
		avatar, err := database.GetUserProfileItem(userID, "image")
		if err != nil {
			log.Printf("Failed to get user profile image: %v\n", err)
			helpers.HTTPError(w, "Failed to get user profile image", http.StatusInternalServerError)
			return
		}
		update.Avatar = avatar.(string)
	}

	updatedProfile := models.UserProfile{
		UserID:         userID,
		NickName:       update.NickName,
		FirstName:      update.FirstName,
		LastName:       update.LastName,
		Gender:         update.Gender,
		DateOfBirth:    update.DateOfBirth,
		Avatar:         update.Avatar,
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

/*
	getProfileByUserName handles the HTTP GET request to retrieve the profile information of a user by their username.
	It extracts the userID from the request context, fetches the userID of the profile being requested,
	and then retrieves and returns the profile data in JSON format.

	URL Path: /api/profile/{user_name}

	Path Parameter:
	- user_name: string - the username of the user whose profile is being requested.

	Response Body:
	{
		"user_name": string,
		"nick_name": string,
		"first_name": string,
		"last_name": string,
		"date_of_birth": "0001-01-01T00:00:00Z",
		"avatar": string,
		"about": string,
		"profile_privacy": "public" | "private",
		"follow_status": "following" | "not_following" | "pending"
		"following_count": int,
		"follower_count": int
	}

	If any error occurs during the process, it returns the corresponding HTTP error status code.
*/func getProfileByUserName(w http.ResponseWriter, r *http.Request) {
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
	
	if UserPageID == userID {
		getMyProfile(w, r)
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
	followingCount, err := database.GetFollowingCount(UserPageID)
	if err != nil {
		log.Printf("Failed to get following count: %v\n", err)
		helpers.HTTPError(w, "Failed to get following count", http.StatusInternalServerError)
		return
	}

	followerCount, err := database.GetFollowerCount(UserPageID)
	if err != nil {
		log.Printf("Failed to get follower count: %v\n", err)
		helpers.HTTPError(w, "Failed to get follower count", http.StatusInternalServerError)
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
		Avatar:         prof.Avatar,
		About:          prof.About,
		ProfilePrivacy: prof.ProfilePrivacy,
		Follow_status:  database.GetFollowStatus(userID, UserPageID),
		FollowingCount: followingCount,
		FollowerCount:  followerCount,
	}
	if err := json.NewEncoder(w).Encode(profile); err != nil {
		log.Printf("Failed to encode profile data: %v\n", err)
		helpers.HTTPError(w, "Something Went Wrong!!", http.StatusInternalServerError)
		return
	}
}

/*
getProfilePosts handles the HTTP GET request to retrieve posts from a user's profile.
It extracts the userID from the request context, fetches the userID of the profile being requested,
and then retrieves and returns the posts data in JSON format.

URL Path: /api/profile/{user_name}/posts

Path Parameter:
- user_name: string - the username of the user whose profile posts are being requested.

Response Body:
[

	{
		"post_id": int,
		"poster": {
				"user_name": string,
				"first_name": string,
				"last_name": string,
				"avatar": string // optional
		},
		"title": string,
		"content": string,
		"image": string, // optional
		"creation_date": "2024-07-07T14:28:45.127591Z", // unix time
		"post_privacy": "private"| "public" | "almost_private",
		"likes_count": int,
		"comments_count": int,
		"likers_usernames": null | []string // optional
		"is_liked": bool
	}

]

If any error occurs during the process, it returns the corresponding HTTP error status code.
*/
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
