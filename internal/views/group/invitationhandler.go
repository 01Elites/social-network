package group

import (
	"encoding/json"
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
	inviteID, err := database.CreateInvite(invite.GroupID, userID, invite.ReceiverID)
	if err != nil {
		helpers.HTTPError(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(inviteID)
	w.WriteHeader(http.StatusOK)
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
	err = database.RespondToInvite(response, userID)
	if err != nil {
		helpers.HTTPError(w, err.Error(), http.StatusNotFound)
		return
	}
	responseIDjson := struct {
		ID int `json:"id"`
	}{
		ID: response.ID,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(responseIDjson)
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
		helpers.HTTPError(w, err.Error(), http.StatusBadRequest)
		return
	}
	requestID, err := database.CreateRequest(request.GroupID, userID)
	if err != nil {
		helpers.HTTPError(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(requestID)
	w.WriteHeader(http.StatusOK)
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
		helpers.HTTPError(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = database.UpdateNotificationTable(response.NotificationID, response.Status, userID)
	if err != nil {
		helpers.HTTPError(w, err.Error(), http.StatusNotFound)
		return
	}
	
	err = database.RespondToRequest(response)
	if err != nil {
		helpers.HTTPError(w, err.Error(), http.StatusNotFound)
		return
	}
	requestIDjson := struct {
		ID int `json:"id"`
	}{
		ID: response.ID,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(requestIDjson)
}
