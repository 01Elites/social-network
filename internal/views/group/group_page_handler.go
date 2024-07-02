package group

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"social-network/internal/database"
	"social-network/internal/models"
	"social-network/internal/views/middleware"
)

func GetGroupPageHandler(w http.ResponseWriter, r *http.Request) {
	var group models.GroupFeed
	var err error
	groupIDstr := r.PathValue("id")
	group.ID, err = strconv.Atoi(groupIDstr)
	if group.ID == 0 || err != nil {
		http.Error(w, "Invalid group ID", http.StatusBadRequest)
		return
	}
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	if group.IsMember, err = database.GroupMember(userID, group.ID); err != nil {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	} 
	if !group.IsMember {
		json.NewEncoder(w).Encode(group)
		w.WriteHeader(http.StatusOK)
		return
	}
	group.Posts, err = database.GetGroupPosts(group.ID)
	if err != nil {
		http.Error(w, "Failed to get group post", http.StatusNotFound)
		return
	}
	group.Members, err = database.GetGroupMembers(group.ID)
	if err != nil {
		http.Error(w, "Failed to get group members", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "get post successful")
	json.NewEncoder(w).Encode(group)
}
