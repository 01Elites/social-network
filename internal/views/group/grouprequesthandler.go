package group

import (
	"encoding/json"
	"log"
	"net/http"

	database "social-network/internal/database/querys"
	"social-network/internal/helpers"
	"social-network/internal/models"
	"social-network/internal/views/middleware"
	"social-network/internal/views/websocket"
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
	if requestMade {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("request already made")
		return
	}
	requestID, groupCreatorID, groupCreator, err := database.CreateRequest(request.GroupID, userID)
	if err != nil {
		helpers.HTTPError(w, "failed to create request", http.StatusNotFound)
		return
	}
	notificationID, err := database.AddToNotificationTable(groupCreatorID, "join_request", requestID)
	if err != nil {
		log.Println("error adding notification to database")
		return
	}
	notification, err := database.GetGroupRequestData(userID, requestID)
	if err != nil {
		log.Println("Failed to get group request")
		helpers.HTTPError(w, "Something Went Wrong with the group Request!!", http.StatusBadRequest)
		return
	}
	notification.ID = notificationID
	websocket.SendNotificationToChannel(*notification, websocket.JoinRequestChan)
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
		"requester":"string" //username
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
	log.Print(response)
	if response.Status != "accepted" && response.Status != "rejected" {
		helpers.HTTPError(w, "status can only be rejected or accepted", http.StatusBadRequest)
		return
	}
	groupExists := database.CheckGroupID(response.GroupID)
	if !groupExists {
		helpers.HTTPError(w, "group ID does not exist", http.StatusBadRequest)
		return
	}
	response.RequesterID, err = database.GetUserIDByUserName(response.Requester)
	if err != nil {
		helpers.HTTPError(w, "error getting user ID", http.StatusBadRequest)
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

	creatorID, err := database.GetGroupCreatorID(response.GroupID)
	if err != nil || creatorID != userID {
		helpers.HTTPError(w, "only group creator can respond to request", http.StatusBadRequest)
		return
	}
	log.Print(response, "22")

	requestID, err := database.RespondToRequest(response)
	if err != nil {
		helpers.HTTPError(w, "error when responding to request", http.StatusNotFound)
		return
	}
	log.Print(response, "22")

	err = database.UpdateNotificationTable(requestID, response.Status, "join_request", userID)
	if err != nil {
		log.Print(err)
		helpers.HTTPError(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
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
		log.Println(err)
		return
	}
	groupExists := database.CheckGroupID(request.GroupID)
	if !groupExists {
		helpers.HTTPError(w, "group ID does not exist", http.StatusBadRequest)
		log.Println("group does not exist")
		return
	}
	isMember, err := database.GroupMember(userID, request.GroupID)
	if err != nil {
		helpers.HTTPError(w, "error checking if user is a member", http.StatusBadRequest)
		log.Println("error checking if user is a member")
		return
	}
	if isMember {
		helpers.HTTPError(w, "you already are a member", http.StatusBadRequest)
		log.Println("user is a member")
		return
	}
	requestID, err := database.CancelRequest(request.GroupID, userID)
	if err != nil {
		helpers.HTTPError(w, "failed to cancel request", http.StatusNotFound)
		log.Println("failed to cancel request")
		return
	}
	creatorID, err := database.GetGroupCreatorID(request.GroupID)
	if err != nil {
		helpers.HTTPError(w, "error getting creator ID", http.StatusNotFound)
		log.Println("error getting creator ID")
		return
	}
	err = database.UpdateNotificationTable(requestID, "canceled", "join_request", creatorID)
	if err != nil {
		helpers.HTTPError(w, err.Error(), http.StatusNotFound)
		log.Println("error updating notification table")
		return
	}
	w.WriteHeader(http.StatusOK)
}
