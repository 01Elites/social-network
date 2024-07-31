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

// GetUsersLiteInfo returns the lite info of the users
func GetUsersLiteInfo(users []string) []models.UserLiteInfo {
	var usersLite []models.UserLiteInfo
	var userLiteInfo models.UserLiteInfo
	for _, user := range users {
		userLite, err := database.GetUserProfileByUserName(user)
		if err != nil {
			log.Printf("Failed to get user lite info: %v\n", err)
			continue
		}
		userLiteInfo.UserName = userLite.Username
		userLiteInfo.FirstName = userLite.FirstName
		userLiteInfo.LastName = userLite.LastName
		userLiteInfo.Avatar = userLite.Avatar
		usersLite = append(usersLite, userLiteInfo)
	}
	return usersLite
}

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

	if len(followers) != 0 {
		friendsList.Followers = GetUsersLiteInfo(followers)
	}

	following, err := database.GetUserFollowingUserNames(userID)
	if err != nil {
		log.Printf("Failed to get following: %v\n", err)
		helpers.HTTPError(w, "Failed to get following", http.StatusInternalServerError)
		return
	}
	if len(following) != 0 {
		friendsList.Following = GetUsersLiteInfo(following)
	}

	// If the user is the same as the user page, return the friend requests and explore
	friend_requests, err := database.GetFollowRequests(userID)
	if err != nil {
		log.Printf("Failed to get friend requests: %v\n", err)
		helpers.HTTPError(w, "Failed to get friend requests", http.StatusInternalServerError)
		return
	}

	// friendsList.Friend_requests = friend_requests
	if len(friend_requests) != 0 {
		for i, request := range friend_requests {
			userLite, err := database.GetUserProfileByUserName(request.UserName)
			if err != nil {
				log.Printf("Failed to get user lite info: %v\n", err)
				continue
			}
			friend_requests[i].UserInfo.UserName = userLite.Username
			friend_requests[i].UserInfo.FirstName = userLite.FirstName
			friend_requests[i].UserInfo.LastName = userLite.LastName
			friend_requests[i].UserInfo.Avatar = userLite.Avatar
		}
		friendsList.Friend_requests = friend_requests
	}

	explore, err := database.GetExplore(userID)
	if err != nil {
		log.Printf("Failed to get explore: %v\n", err)
		helpers.HTTPError(w, "Failed to get explore", http.StatusInternalServerError)
		return
	}
	if len(explore) != 0 {
		friendsList.Explore = GetUsersLiteInfo(explore)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(friendsList)
}

// GetFriendsHandler returns the friends of the user
func GetFriendsHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the userID from context using the same key defined globally
	_, ok := r.Context().Value(middleware.UserIDKey).(string)
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

	if len(followers) != 0 {
		friendsList.Followers = GetUsersLiteInfo(followers)
	}

	following, err := database.GetUserFollowingUserNames(UserPageID)
	if err != nil {
		log.Printf("Failed to get following: %v\n", err)
		helpers.HTTPError(w, "Failed to get following", http.StatusInternalServerError)
		return
	}

	if len(following) != 0 {
		friendsList.Following = GetUsersLiteInfo(following)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(friendsList)
}
