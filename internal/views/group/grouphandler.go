package group

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	database "social-network/internal/database/querys"
	"social-network/internal/helpers"
	"social-network/internal/models"
	"social-network/internal/views/middleware"
)

/*
CreateGroupHandler creates a new group.

This function creates a new group using the data provided in the request body.
It requires a valid user session to create a group.

Example:

	    // To create a new group
	    POST /api/group
	   Body: {
			"title":"string",
			"description":"string"
			}
*/
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

/*
GetGroupPageHandler retrieves a group page by its ID.
This function retrieves a group based on the provided group ID.
It requires a valid user session to access the group.
The page defers depending on whether the user a part of the group.

Example:

	// To retrieve a group with ID 123
	GET api/group/123

Response:

	    {
	    "id": 0,
	    "members":"[]string",
	    "posts":"[]posts",
	    "ismember":false
	}
*/
func GetGroupPageHandler(w http.ResponseWriter, r *http.Request) {
	var group models.GroupFeed
	var err error
	groupIDstr := r.PathValue("id")
	group.ID, err = strconv.Atoi(groupIDstr)
	groupExists := database.CheckGroupID(group.ID) // check if the group has been created
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
	if !group.IsMember { // to view group page for a non-member
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


/*
ExitGroupHandler exits the group which has the
group ID provided sent in the request body
	    // To create a new group
	    POST /api/group
	   Body: {
			"id":0
			}
*/
func ExitGroupHandler(w http.ResponseWriter, r *http.Request) {
	var group models.ID
	err := json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		helpers.HTTPError(w, err.Error(), http.StatusBadRequest)
		return
	}
	groupExists := database.CheckGroupID(group.ID) // check if the group has been created
	if !groupExists {
		http.Error(w, "Invalid group ID", http.StatusBadRequest)
		return
	}
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	isMember, err := database.GroupMember(userID, group.ID)
		if err != nil {
		helpers.HTTPError(w, "Error when checking if user is a member", http.StatusBadRequest)
		return
	}
	if !isMember { // to view group page for a non-member
		helpers.HTTPError(w, "user not part of group", http.StatusBadRequest)
		return
	}
	if err := database.LeaveGroup(userID, group.ID);err != nil {
		helpers.HTTPError(w, "error leaving group", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(group)
}
