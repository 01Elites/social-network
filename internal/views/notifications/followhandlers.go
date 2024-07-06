package notifications

import (
	"encoding/json"
	"net/http"
	database "social-network/internal/database/querys"
	"social-network/internal/models"
	"social-network/internal/views/middleware"
)

// FollowHandler creates a follow request for a user. It expects a JSON body with the following format:
//
//	{
//		"receiver": "username" // the username of the user to follow
//	}
func FollowHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	var request models.Request
	request.Sender = userID
	// Decode the request body into the request struct
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonError := models.Error{
			Reason: err.Error(),
		}
		json.NewEncoder(w).Encode(jsonError)
		return
	}
	// Get the userID of the receiver
	request.Receiver, err = database.GetUserIDByUserName(request.Receiver)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonError := models.Error{
			Reason: err.Error(),
		}
		json.NewEncoder(w).Encode(jsonError)
		return
	}
	// Check if the user is trying to follow themselves
	if request.Receiver == userID {
		w.WriteHeader(http.StatusBadRequest)
		jsonError := models.Error{
			Reason: "You can't follow yourself",
		}
		json.NewEncoder(w).Encode(jsonError)
		return
	}

	// Check if the user is already following the receiver
	isFollowing := database.IsFollowing(request.Sender, request.Receiver)

	// already following , unfollow the user
	if isFollowing {
		err = database.UnFollowUser(request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			jsonError := models.Error{
				Reason: err.Error(),
			}
			json.NewEncoder(w).Encode(jsonError)
			return
		}
		// Notifications(request , "unfollow")

		w.WriteHeader(http.StatusOK)
		return
	}

	// Check if receiver account is private
	ReceiverProf, err := database.GetUserProfile(request.Receiver)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonError := models.Error{
			Reason: err.Error(),
		}
		json.NewEncoder(w).Encode(jsonError)
		return
	}
	// Check if the receiver account is private or public
	if ReceiverProf.ProfilePrivacy == "private" {
		request.ID, err = database.CreateFollowRequest(request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			jsonError := models.Error{
				Reason: err.Error(),
			}
			json.NewEncoder(w).Encode(jsonError)
			return
		}
		// Notifications(request , "followRequest")

	} else {
		err = database.FollowUser(request)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			jsonError := models.Error{
				Reason: err.Error(),
			}
			json.NewEncoder(w).Encode(jsonError)
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
//		"status": "accept" or "reject" // the status of the follow request
//	}
func RespondToFollowHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	var response models.Response
	response.Followee = userID
	// Check if Followee account is private
	FolloweeProf, err := database.GetUserProfile(response.Followee)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonError := models.Error{
			Reason: err.Error(),
		}
		json.NewEncoder(w).Encode(jsonError)
		return
	}
	
	// Check if the receiver account is private or public
	if FolloweeProf.ProfilePrivacy != "private" {
		w.WriteHeader(http.StatusBadRequest)
		jsonError := models.Error{
			Reason: "Bro Your account is public, you don't have any follow request",
		}
		json.NewEncoder(w).Encode(jsonError)
		return
	}

	// Decode the request body into the response struct
	err = json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonError := models.Error{
			Reason: err.Error(),
		}
		json.NewEncoder(w).Encode(jsonError)
		return
	}
	// Get the userID of the follower
	response.Follower, err = database.GetUserIDByUserName(response.Follower)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonError := models.Error{
			Reason: err.Error(),
		}
		json.NewEncoder(w).Encode(jsonError)
		return
	}

	err = database.RespondToFollow(response)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonError := models.Error{
			Reason: err.Error(),
		}
		json.NewEncoder(w).Encode(jsonError)
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
