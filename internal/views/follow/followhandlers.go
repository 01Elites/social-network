package follow

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	database "social-network/internal/database/querys"
	"social-network/internal/helpers"
	"social-network/internal/models"
	"social-network/internal/views/middleware"
	"social-network/internal/views/websocket"
)

// FollowHandler creates a follow request for a user. It expects a JSON body with the following format:
//
//	{
//		"receiver": "username" // the username of the user to follow
//	}
func FollowHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		log.Printf("Error extracting token. User ID not found\n")
		helpers.HTTPError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var request models.Request
	request.Sender = userID
	// Decode the request body into the request struct
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Printf("Error decoding follow request: %v\n", err)
		helpers.HTTPError(w, "Something Went Wrong!!", http.StatusBadRequest)
		return
	}
	// Get the userID of the receiver
	request.Receiver, err = database.GetUserIDByUserName(request.Receiver)
	if err != nil {
		log.Printf("Error while gtting UserID By UserName: %v\n", err)
		helpers.HTTPError(w, "Something Went Wrong!!", http.StatusBadRequest)
		return
	}
	// Check if the user is trying to follow themselves
	if request.Receiver == userID {
		helpers.HTTPError(w, "You can't follow yourself!!", http.StatusBadRequest)
		return
	}

	// Check if the user is already following the receiver
	isFollowing := database.IsFollowing(request.Sender, request.Receiver)

	// already following , unfollow the user
	if isFollowing {
		err = database.UnFollowUser(request)
		if err != nil {
			helpers.HTTPError(w, "Something Went Wrong with unfollow!!", http.StatusBadRequest)
			return
		}
		// Notifications(request , "unfollow")

		w.WriteHeader(http.StatusOK)
		return
	}
	// Check if receiver account is private
	isPrivateReceiver, err := database.IsPrivateUser(request.Receiver)
	if err != nil {
		helpers.HTTPError(w, "Something is Wrong!!", http.StatusBadRequest)
		return
	}
	// Check if the receiver account is private or public
	if !isPrivateReceiver {
		err = database.FollowUser(request)
		if err != nil {
			helpers.HTTPError(w, "Something Went Wrong with the Following Request!!", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}
	// If the receiver account is private, create a follow request
	if err := database.CreateFollowRequest(&request); err != nil {
		helpers.HTTPError(w, "Something is Wrong with the Follow Request!!", http.StatusBadRequest)
		return
	}
	if request.Status != "canceled" {
		notificationID, err := database.AddToNotificationTable(request.Receiver, "follow_request", request.ID)
		if err != nil {
			log.Println("error adding notification to database:", err)
			helpers.HTTPError(w, "Something Went Wrong with the Follow Request!!", http.StatusBadRequest)
			return
		}
		recieverUsername, err := database.GetUserNameByID(request.Receiver)
		if err != nil {
			log.Println("Error getting user name:", err)
			helpers.HTTPError(w, "Something Went Wrong with the Follow Request!!", http.StatusBadRequest)
			return
		}
		recieverConnected := websocket.IsUserConnected(recieverUsername)
		if !recieverConnected {
			w.WriteHeader(http.StatusOK)
			return
		}
		notification, err := database.GetFollowRequestNotification(request)
		if err != nil {
			log.Println("Failed to get follow request")
			helpers.HTTPError(w, "Something Went Wrong with the Follow Request!!", http.StatusBadRequest)
			return
		}
		notification.ID = notificationID
		fmt.Println("hello before")
		websocket.SendNotificationToChannel(*notification, websocket.FollowRequestChan)
		fmt.Println("hello after")
	}

	w.WriteHeader(http.StatusOK)
}

// RespondToFollowHandler responds to a follow request. It expects a JSON body with the following format:
//
//	{
//		"follower": "username", // the username of the user who sent the follow request
//		"status": "accepted" or "rejected" // the status of the follow request
//	}
func RespondToFollowHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		log.Printf("Error extracting token. User ID not found\n")
		helpers.HTTPError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	var response models.Response
	response.Followee = userID
	// Check if Followee account is private
	isPrivateFollowee, err := database.IsPrivateUser(response.Followee)
	if err != nil {
		helpers.HTTPError(w, "Something is Wrong!!", http.StatusBadRequest)
		return
	}

	// Check if the receiver account is private or public
	if !isPrivateFollowee {
		helpers.HTTPError(w, "Bro Your account is public, you don't have any follow request!!", http.StatusBadRequest)
		return
	}

	// Decode the request body into the response struct
	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		log.Printf("Error decoding follow Response: %v\n", err)
		helpers.HTTPError(w, "Something Went Wrong!!", http.StatusBadRequest)
		return
	}
	// Get the userID of the follower
	response.Follower, err = database.GetUserIDByUserName(response.Follower)
	if err != nil {
		log.Printf("Error while gtting UserID By UserName: %v\n", err)
		helpers.HTTPError(w, "Something Went Wrong!!", http.StatusBadRequest)
		return
	}

	err = database.RespondToFollow(&response)
	if err != nil {
		helpers.HTTPError(w, "Something Went Wrong with the Follow Response!!", http.StatusBadRequest)
		return
	}

	// Update the notification table
	err = database.UpdateNotificationTable(response.ID, response.Status, "follow_request", userID)
	if err != nil {
		helpers.HTTPError(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
