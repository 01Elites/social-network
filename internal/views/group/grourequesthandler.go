package group

import (
	"encoding/json"
	"net/http"

	database "social-network/internal/database/querys"
	"social-network/internal/helpers"
	"social-network/internal/models"
	"social-network/internal/views/middleware"
	"social-network/internal/views/websocket"
	"social-network/internal/views/websocket/types"
)

/*
CreateRequestHandler creates a request to a certain group.
This function creates a new request using the groupID
provided in the request body.
It requires a valid user session and the user should not
be a member to create a request.

Example:

	// To create a new invite

/api/grouprequest

	Body:{
	"group_id":0
	}
*/
func CreateRequestHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	var request models.GroupAction
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		helpers.HTTPError(w, "failed to decode request", http.StatusBadRequest)
		return
	}
	groupExists := database.CheckGroupID(request.GroupID)
	if !groupExists {
		helpers.HTTPError(w, "group ID does not exist", http.StatusBadRequest)
		return
	}

	isMember, err := database.GroupMember(userID, request.GroupID)
	if err != nil {
		helpers.HTTPError(w, "error checking if user is a member", http.StatusBadRequest)
		return
	}
	if isMember {
		helpers.HTTPError(w, "you already are a member", http.StatusBadRequest)
		return
	}
	requestMade, err := database.CheckForGroupRequest(request.GroupID, userID)
	if err != nil {
		helpers.HTTPError(w, "failed to create request", http.StatusNotFound)
		return
	}
	if requestMade{
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("request already made")
		return
	}
	groupCreatorID, err := database.CreateRequest(request.GroupID, userID)
	if err != nil {
		helpers.HTTPError(w, "failed to create request", http.StatusNotFound)
		return
	}
	groupCreator,err := database.GetUserNameByID(groupCreatorID)
	if err != nil {
		helpers.HTTPError(w, "failed to get Username", http.StatusNotFound)
		return
	}
	event := types.Event{
		Type:    "JoinRequest",
		ToUser:  groupCreator,
		Payload: request, // You can customize the payload as per your requirements
	}
	websocket.JoinRequestChan <- event
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(groupCreator)
}

/*
RequestResponseHandler responds to a group join request.
This function responds to a request using the status
provided in the request body.
It requires a valid user session and the user should be
the creator of the group to respond to the request.

Example:

	 // To respond to a request
	POST /api/groupresponse
		Body:{
		"group_id":0
		"response":"accepted" | "rejected"
		}
*/
func RequestResponseHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	var response models.GroupResponse
	err := json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		helpers.HTTPError(w, "error decoding response", http.StatusBadRequest)
		return
	}
	if response.Status != "accepted" && response.Status != "rejected" {
		helpers.HTTPError(w, "status can only be rejected or accepted", http.StatusBadRequest)
		return
	}
	groupExists := database.CheckGroupID(response.GroupID)
	if !groupExists {
		helpers.HTTPError(w, "group ID does not exist", http.StatusBadRequest)
		return
	}
	isMember, err := database.GroupMember(response.RequesterID, response.GroupID)
	if err != nil {
		helpers.HTTPError(w, "error checking if user is a member", http.StatusBadRequest)
		return
	}
	if isMember {
		helpers.HTTPError(w, "user already a member", http.StatusBadRequest)
		return
	}
	isCreator := database.CheckGroupCreator(userID, response.GroupID)
	if !isCreator {
		helpers.HTTPError(w, "only group creator can respond to request", http.StatusBadRequest)
		return
	}
	err = database.RespondToRequest(response)
	if err != nil {
		helpers.HTTPError(w, "error when responding to request", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	// err = database.UpdateNotificationTable(response.NotificationID, response.Status, userID)
	// if err != nil {
	// 	helpers.HTTPError(w, err.Error(), http.StatusNotFound)
	// 	return
	// }
}

func CancelRequestHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	var request models.GroupAction
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		helpers.HTTPError(w, "failed to decode request", http.StatusBadRequest)
		return
	}
	groupExists := database.CheckGroupID(request.GroupID)
	if !groupExists {
		helpers.HTTPError(w, "group ID does not exist", http.StatusBadRequest)
		return
	}
	isMember, err := database.GroupMember(userID, request.GroupID)
	if err != nil {
		helpers.HTTPError(w, "error checking if user is a member", http.StatusBadRequest)
		return
	}
	if isMember {
		helpers.HTTPError(w, "you already a member", http.StatusBadRequest)
		return
	}
	err = database.CancelRequest(request.GroupID, userID)
	if err != nil {
		helpers.HTTPError(w, "failed to cancel request", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
