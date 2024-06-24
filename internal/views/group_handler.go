package views

import (
	"encoding/json"
	"net/http"
	"social-network/internal/models"
	"social-network/internal/database"
)

func CreateGroupHandler(w http.ResponseWriter, r *http.Request){
	var user models.User
		// user := ValidateSession(w, r)
	// if user == nil {
	// 	return
	// }
	var group models.Create_Group
	err := json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		http.Error(w, "Failed to decode group", http.StatusBadRequest)
		return
	}
	err = database.CreateGroup(user.UserID, group)
	if err != nil {
		http.Error(w, "Failed to create group", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}