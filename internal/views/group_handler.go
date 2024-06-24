package views

import (
	"encoding/json"
	"net/http"
	"social-network/internal/events"
	"social-network/internal/database"
)

func CreateGroupHandler(w http.ResponseWriter, r *http.Request){
	var user events.User
		// user := ValidateSession(w, r)
	// if user == nil {
	// 	return
	// }
	var group events.Create_Group
	err := json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		http.Error(w, "Failed to decode group", http.StatusBadRequest)
		return
	}
	err = database.CreateGroup(user.ID, group)
	if err != nil {
		http.Error(w, "Failed to create group", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}