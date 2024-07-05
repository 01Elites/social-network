package group

import (
	"encoding/json"
	"net/http"
	database "social-network/internal/database/querys"
	"social-network/internal/models"
	"social-network/internal/views/middleware"
)

func CreateInvitationHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	var invite models.GroupInvite
	err := json.NewDecoder(r.Body).Decode(&invite)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonError := models.Error{
			Reason: err.Error(),
		}
		json.NewEncoder(w).Encode(jsonError)
		return
	}
	inviteID, err := database.CreateInvite(invite.GroupID, userID, invite.ReceiverID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		jsonError := models.Error{
			Reason: "invalid post id",
		}
		json.NewEncoder(w).Encode(jsonError)
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
	var response models.InviteResponse
	err := json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonError := models.Error{
			Reason: err.Error(),
		}
		json.NewEncoder(w).Encode(jsonError)
		return
	}
	err = database.RespondToInvite(response, userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonError := models.Error{
			Reason: err.Error(),
		}
		json.NewEncoder(w).Encode(jsonError)
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
