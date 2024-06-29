package views

import (
	"encoding/json"
	"net/http"
	"social-network/internal/models"
	"social-network/internal/database"
	"strconv"
)

func CreateInvitationHandler(w http.ResponseWriter, r *http.Request){
	var to string
	from, ok := r.Context().Value(userIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	groupIDstr := r.PathValue("id")
	groupID, err := strconv.Atoi(groupIDstr)
	if groupID == 0 || err != nil{
		w.WriteHeader(http.StatusBadRequest)
		jsonError := models.Error{
				Reason: "invalid group id",
		}
		json.NewEncoder(w).Encode(jsonError)
		return
	}
	err = json.NewDecoder(r.Body).Decode(&to)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		jsonError := models.Error{
				Reason: "invalid request",
		}
		json.NewEncoder(w).Encode(jsonError)
		return
	}
	if err != nil {
		http.Error(w, "Failed to decode post", http.StatusBadRequest)
		return
	}
	err = database.CreateInvite(groupID, from, to)
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