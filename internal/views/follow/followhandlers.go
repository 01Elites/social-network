package follow

import (
	"encoding/json"
	"log"
	"net/http"

	database "social-network/internal/database/querys"
	"social-network/internal/helpers"
	"social-network/internal/models"
	"social-network/internal/views/middleware"
	"social-network/internal/views/websocket"
	"social-network/internal/views/websocket/types"
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
	if isPrivateReceiver {
		request.ID, err = database.CreateFollowRequest(request)
		if err != nil {
			helpers.HTTPError(w, "Something is Wrong with the Follow Request!!", http.StatusBadRequest)
			return
		}
		// Notifications(request , "followRequest")
		// Here, insert in Channel1
		event := types.Event{
			Type:    "follow request",
			Payload: request, // You can customize the payload as per your requirements
		}
		websocket.Channel1 <- event

	} else {
		err = database.FollowUser(request)
		if err != nil {
			helpers.HTTPError(w, "Something Went Wrong with the Following Request!!", http.StatusBadRequest)
			return
		}
		// Notifications(request , "follow")

	}
	// requestIDjson := struct {
	// 	ID int `json:"id"`
	// }{
	// 	ID: request.ID,
	// }

	// Notifications(request , "follow")

	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(requestIDjson)
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
	if isPrivateFollowee {
		helpers.HTTPError(w, "Bro Your account is public, you don't have any follow request!!", http.StatusBadRequest)
		return
	}

	// Decode the request body into the response struct
	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		log.Printf("Error decoding follow Respons: %v\n", err)
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

	err = database.RespondToFollow(response)
	if err != nil {
		helpers.HTTPError(w, "Something Went Wrong with the Follow Respons!!", http.StatusBadRequest)
		return
	}
	// requestIDjson := struct {
	// 	ID int `json:"id"`
	// }{
	// 	ID: response.ID,
	// }
	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(requestIDjson)
}
