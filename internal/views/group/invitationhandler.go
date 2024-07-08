package group

import (
	"encoding/json"
	"log"
	"net/http"
	database "social-network/internal/database/querys"
	"social-network/internal/helpers"
	"social-network/internal/models"
	"social-network/internal/views/middleware"
)

func CreateInvitationHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	var invite models.GroupAction
	err := json.NewDecoder(r.Body).Decode(&invite)
	if err != nil {
		helpers.HTTPError(w, err.Error(), http.StatusBadRequest)
		return
	}
	groupExists := database.CheckGroupID(invite.GroupID)
	if !groupExists {
		helpers.HTTPError(w, "group ID does not exist", http.StatusBadRequest)
		return
	}
	invite.ReceiverID, err = database.GetUserIDByUserName(invite.Username)
	if err != nil {
		helpers.HTTPError(w, "failed to get user ID", http.StatusBadRequest)
		return
	}
	isMember, err := database.GroupMember(userID, invite.GroupID)
	if err != nil {
		helpers.HTTPError(w, "check if user is a member error", http.StatusBadRequest)
		return
	}
	if !isMember {
		helpers.HTTPError(w, "user not a member to make an invitation", http.StatusBadRequest)
		return
	}
	isMember, err = database.GroupMember(invite.ReceiverID, invite.GroupID)
	if err != nil {
		helpers.HTTPError(w, "error checking if user is a member", http.StatusBadRequest)
		return
	}
	if isMember {
		helpers.HTTPError(w, "user already a member", http.StatusBadRequest)
		return
	}
	inviteID, err := database.CreateInvite(invite.GroupID, userID, invite.ReceiverID)
	if err != nil {
		helpers.HTTPError(w, "Failed to create invitation", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Print(inviteID)
	// database.AddToNotificationTable(inviteID)
}

func InvitationResponseHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	var response models.GroupResponse
	err := json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		helpers.HTTPError(w, err.Error(), http.StatusBadRequest)
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

	isMember, err := database.GroupMember(userID, response.GroupID)
	if err != nil {
		helpers.HTTPError(w, "error checking if user is a member", http.StatusBadRequest)
		return
	}
	if isMember {
		helpers.HTTPError(w, "user already a member", http.StatusBadRequest)
		return
	}
	err = database.RespondToInvite(response, userID)
	if err != nil {
		helpers.HTTPError(w, "Failed to respond to invite", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	// database.UpdateNotificationTable()
}

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
		helpers.HTTPError(w, "you already a member", http.StatusBadRequest)
		return
	}
	requestID, err := database.CreateRequest(request.GroupID, userID)
	if err != nil {
		helpers.HTTPError(w, "failed to create request", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(requestID)
}

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
