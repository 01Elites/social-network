package group

import (
	"encoding/json"
	"net/http"
	database "social-network/internal/database/querys"
	"social-network/internal/helpers"
	"social-network/internal/models"
	"social-network/internal/views/middleware"
	"strings"
	"strconv"
)

func CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	var group models.CreateGroup
	err := json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		helpers.HTTPError(w, err.Error(), http.StatusBadRequest)
		return
	}
	group.Title = strings.TrimSpace(group.Title)
	group.Description = strings.TrimSpace(group.Description)
	if group.Title == "" || group.Description == "" {
		helpers.HTTPError(w, "title or description cannot be empty", http.StatusBadRequest)
		return
	}
	groupID, err := database.CreateGroup(userID, group)
	if err != nil {
		helpers.HTTPError(w, "failed to create group", http.StatusBadRequest)
		return
	}
	groupIDjson := struct {
		ID int `json:"id"`
	}{
		ID: groupID,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(groupIDjson)
}

func GetGroupPageHandler(w http.ResponseWriter, r *http.Request) {
	var group models.GroupFeed
	var err error
	groupIDstr := r.PathValue("id")
	group.ID, err = strconv.Atoi(groupIDstr)
	groupExists := database.CheckGroupID(group.ID)
	if group.ID == 0 || err != nil || !groupExists {
		http.Error(w, "Invalid group ID", http.StatusBadRequest)
		return
	}
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	if group.IsMember, err = database.GroupMember(userID, group.ID); err != nil {
		helpers.HTTPError(w, "Error when checking if user is a member", http.StatusBadRequest)
		return
	}
	if !group.IsMember {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(group)
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
	json.NewEncoder(w).Encode(group)
}