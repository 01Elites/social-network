package group

import (
	"encoding/json"
	"net/http"
	database "social-network/internal/database/querys"
	"social-network/internal/helpers"
	"social-network/internal/models"
	"social-network/internal/views/middleware"
)

func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	var request models.GroupEvent
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
	if !isMember {
		helpers.HTTPError(w, "you are not a memeber", http.StatusBadRequest)
		return
	}

	err = database.CreateEvent(request.GroupID, userID, request.Title, request.Description, request.EventTime)
	if err != nil {
		helpers.HTTPError(w, "failed to cancel request", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}
