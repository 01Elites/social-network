package friends

import (
	"encoding/json"
	"log"
	"net/http"
	database "social-network/internal/database/querys"
	"social-network/internal/helpers"
	"social-network/internal/models"
	"social-network/internal/views/middleware"
	"strings"
)

// GetMyFriendsHandler returns the friends of the user
func GetMyFriendsHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the userID from context using the same key defined globally
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		log.Printf("User ID not found")
		helpers.HTTPError(w, "User ID not found", http.StatusInternalServerError)
		return
	}

	var friendsList models.Friends

	followers, err := database.GetUserFollowerUserNames(userID)
	if err != nil {
		log.Printf("Failed to get followers: %v\n", err)
		helpers.HTTPError(w, "Failed to get followers", http.StatusInternalServerError)
		return
	}
	friendsList.Followers = followers

	following, err := database.GetUserFollowingUserNames(userID)
	if err != nil {
		log.Printf("Failed to get following: %v\n", err)
		helpers.HTTPError(w, "Failed to get following", http.StatusInternalServerError)
		return
	}
	friendsList.Following = following

	// If the user is the same as the user page, return the friend requests and explore
	friend_requests, err := database.GetFollowRequests(userID)
	if err != nil {
		log.Printf("Failed to get friend requests: %v\n", err)
		helpers.HTTPError(w, "Failed to get friend requests", http.StatusInternalServerError)
		return
	}
	friendsList.Friend_requests = friend_requests

	explore, err := database.GetExplore(userID)
	if err != nil {
		log.Printf("Failed to get explore: %v\n", err)
		helpers.HTTPError(w, "Failed to get explore", http.StatusInternalServerError)
		return
	}
	friendsList.Explore = explore

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(friendsList)
}

// GetFriendsHandler returns the friends of the user
func GetFriendsHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the userID from context using the same key defined globally
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		log.Printf("User ID not found")
		helpers.HTTPError(w, "User ID not found", http.StatusInternalServerError)
		return
	}

	userName := strings.TrimPrefix(r.URL.Path, "/api/friends/")
	userName = strings.TrimSuffix(userName, "/")

	if userName == "" {
		helpers.HTTPError(w, "username cannot be empty", http.StatusBadRequest)
		return
	}
	UserPageID, err := database.GetUserIDByUserName(userName)
	if err != nil {
		log.Printf("Failed to get userID: %v\n", err)
		helpers.HTTPError(w, "Failed to get userID", http.StatusInternalServerError)
		return
	}

	var friendsList models.Friends
	friendsList.UserName = userName

	followers, err := database.GetUserFollowerUserNames(UserPageID)
	if err != nil {
		log.Printf("Failed to get followers: %v\n", err)
		helpers.HTTPError(w, "Failed to get followers", http.StatusInternalServerError)
		return
	}
	friendsList.Followers = followers

	following, err := database.GetUserFollowingUserNames(UserPageID)
	if err != nil {
		log.Printf("Failed to get following: %v\n", err)
		helpers.HTTPError(w, "Failed to get following", http.StatusInternalServerError)
		return
	}
	friendsList.Following = following

	// If the user is the same as the user page, return the friend requests and explore
	if userID == UserPageID {
		friend_requests, err := database.GetFollowRequests(UserPageID)
		if err != nil {
			log.Printf("Failed to get friend requests: %v\n", err)
			helpers.HTTPError(w, "Failed to get friend requests", http.StatusInternalServerError)
			return
		}
		friendsList.Friend_requests = friend_requests

		explore, err := database.GetExplore(UserPageID)
		if err != nil {
			log.Printf("Failed to get explore: %v\n", err)
			helpers.HTTPError(w, "Failed to get explore", http.StatusInternalServerError)
			return
		}
		friendsList.Explore = explore
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(friendsList)
}
