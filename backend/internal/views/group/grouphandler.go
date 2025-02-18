package group

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

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
	if len(group.Title) >13 || len(group.Description) > 200 {
		helpers.HTTPError(w, "title or description is too long", http.StatusBadRequest)
		return
	}

	groupID, err := database.CreateGroup(userID, group)
	if err != nil {
		helpers.HTTPError(w, "failed to create group", http.StatusBadRequest)
		return
	}
	err = database.CreatGroupChat(groupID)
	if err != nil {
		helpers.HTTPError(w, "failed to create group chat", http.StatusBadRequest)
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

	       {"id":1,
					"title":"testing",
					"description":"desc test",
					"members":[]string //usernames,
					"ismember":true,
					"request_made":true
	      	"events":[{"id":1,
	      	    "title":"testing",
	            "description":"desc test",
	            "options":["Going","Notgoing"],
	            "event_time":"2024-07-10T14:00:00Z",
	        "responded_users":["AHeidenreich5716","Going"]
	        "creator":{
					"user_name": string,
					"first_name": string,
					"last_name": string,
					"avatar": string // optional
			},
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
	username, err := database.GetUserNameByID(userID)
	if err != nil {
		http.Error(w, "Failed to get username", http.StatusInternalServerError)
		return
	}

	group.Title, group.Description, err = database.GetGroupInfo(group.ID)
	if err != nil {
		http.Error(w, "Failed to get group info", http.StatusInternalServerError)
		return
	}

	group.Members,_, err = database.GetGroupMembers(userID, group.ID)
	if err != nil {
		http.Error(w, "Failed to get group members", http.StatusInternalServerError)
		return
	}

	group.Creator, err = database.GetCreatorProfile(group.ID)
	if err != nil {
		http.Error(w, "Failed to get creator profile", http.StatusInternalServerError)
		return
	}
	if group.Creator.UserName == username {
		group.IsCreator = true
		group.Requesters, err = database.GetGroupRequests(group.ID)
		if err != nil {
			http.Error(w, "Failed to get group requests", http.StatusInternalServerError)
			return
		}
	}
	if group.IsMember, err = database.GroupMember(userID, group.ID); err != nil {
		helpers.HTTPError(w, "Error when checking if user is a member", http.StatusBadRequest)
		return
	}
	if !group.IsMember { // to view group page for a non-member
		group.RequestMade, err = database.CheckForGroupRequest(group.ID, userID)
		if err != nil {
			helpers.HTTPError(w, "failed to check for request", http.StatusNotFound)
			return
		}
		group.InvitedBy, err = database.CheckForGroupInvitation(group.ID, userID)
		if err != nil {
			helpers.HTTPError(w, "failed to check for invitation", http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(group)
		return
	}
	
	group.Explore, err = database.GetExploreGroup(group.ID)
	if err != nil {
		http.Error(w, "Failed to get group explore", http.StatusInternalServerError)
		return
	}

	// group.Events, err = database.GetGroupEvents(group.ID)
	// if err != nil {
	// 	http.Error(w, "Failed to get group events", http.StatusInternalServerError)
	// 	return
	// }
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
	log.Print(group)
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
	creatorID, err := database.GetGroupCreatorID(group.ID)
	if err != nil {
		helpers.HTTPError(w, "Error when checking if user is a member", http.StatusBadRequest)
		return
	}
	if !isMember { // to view group page for a non-member
		helpers.HTTPError(w, "user not part of group", http.StatusBadRequest)
		return
	}
	if creatorID == userID {
		helpers.HTTPError(w, "creator cannot leave group", http.StatusBadRequest)
		return
	}
	if err := database.LeaveGroup(userID, group.ID); err != nil {
		helpers.HTTPError(w, "error leaving group", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(group)
}

func getGroupPostsHandler(w http.ResponseWriter, r *http.Request) {
	groupIDstr := r.PathValue("id")
	groupID, err := strconv.Atoi(groupIDstr)
	groupExists := database.CheckGroupID(groupID) // check if the group has been created
	if groupID == 0 || err != nil || !groupExists {
		http.Error(w, "Invalid group ID", http.StatusBadRequest)
		return
	}
	posts, err := database.GetGroupPosts(groupID)
	if err != nil {
		helpers.HTTPError(w, "Failed to get group posts", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func GetEventsHandler(w http.ResponseWriter, r *http.Request) {
	groupIDstr := r.PathValue("id")
	groupID, err := strconv.Atoi(groupIDstr)
	groupExists := database.CheckGroupID(groupID) // check if the group has been created
	if groupID == 0 || err != nil || !groupExists {
		http.Error(w, "Invalid group ID", http.StatusBadRequest)
		return
	}
	events, err := database.GetGroupEvents(groupID)
	if err != nil {
		helpers.HTTPError(w, "Failed to get group events", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	events = helpers.ArrangeEvents(events)
	var arrangedEvents []models.Event
	for _, event := range events {
		if event.EventTime.After(time.Now()){
			arrangedEvents = append(arrangedEvents, event)
		}
	}
	for _, event := range events {
		if event.EventTime.Before(time.Now()){
			arrangedEvents = append(arrangedEvents, event)
		}
	}
	json.NewEncoder(w).Encode(arrangedEvents)
}