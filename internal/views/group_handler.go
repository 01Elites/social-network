package views

import (
	"encoding/json"
	"net/http"
	"social-network/internal/database"
	"social-network/internal/models"
	"social-network/internal/views/middleware"
)

func CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	var group models.Create_Group
	err := json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonError := models.Error{
				Reason: err.Error(),
		}
		json.NewEncoder(w).Encode(jsonError)
		return
	}
	groupID, err := database.CreateGroup(userID, group)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonError := models.Error{
				Reason: err.Error(),
		}
		json.NewEncoder(w).Encode(jsonError)
			return
	}
	groupIDjson := struct {
		ID int `json:"id"`
	}{
		ID: groupID,
	}
	json.NewEncoder(w).Encode(groupIDjson)
	w.WriteHeader(http.StatusCreated)
}

