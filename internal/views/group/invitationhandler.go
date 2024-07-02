package group

import (
	"encoding/json"
	"net/http"
	"strconv"

	"social-network/internal/database"
	"social-network/internal/models"
	"social-network/internal/views/middleware"
)

func CreateInvitationHandler(w http.ResponseWriter, r *http.Request) {
	fromUser, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	var toUser string
	groupIDstr := r.PathValue("id")
	groupID, err := strconv.Atoi(groupIDstr)
	if groupID == 0 || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonError := models.Error{
			Reason: "invalid group id",
		}
		json.NewEncoder(w).Encode(jsonError)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&toUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonError := models.Error{
			Reason: "invalid request",
		}
		json.NewEncoder(w).Encode(jsonError)
		return
	}
	err = database.CreateInvite(groupID, fromUser, toUser)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		jsonError := models.Error{
			Reason: "invalid post id",
		}
		json.NewEncoder(w).Encode(jsonError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
