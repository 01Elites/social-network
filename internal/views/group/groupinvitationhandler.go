package group

import (
	"encoding/json"
	"log"
	"net/http"
	database "social-network/internal/database/querys"
	"social-network/internal/helpers"
	"social-network/internal/models"
	"social-network/internal/views/middleware"
)

/*
CreateInvitationHandler creates an invitation to a certain group.
This function creates a new invite using the groupID and username
provided in the request body.
It requires a valid user session and the user should be a member
to create an invite.

Example:
	 // To create a new invite
	 POST /api/invitation
		Body:{
			"receiver":"string"  //username
			"group_id":0
			}
*/
func CreateInvitationHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	var invite models.GroupAction
	err := json.NewDecoder(r.Body).Decode(&invite)
	if err != nil {
		helpers.HTTPError(w, err.Error(), http.StatusBadRequest)
		return
	}
	groupExists := database.CheckGroupID(invite.GroupID)
	if !groupExists {
		helpers.HTTPError(w, "group ID does not exist", http.StatusBadRequest)
		return
	}
	invite.ReceiverID, err = database.GetUserIDByUserName(invite.Username)
	if err != nil {
		helpers.HTTPError(w, "failed to get user ID", http.StatusBadRequest)
		return
	}
	isMember, err := database.GroupMember(userID, invite.GroupID)
	if err != nil {
		helpers.HTTPError(w, "check if user is a member error", http.StatusBadRequest)
		return
	}
	if !isMember {
		helpers.HTTPError(w, "user not a member to make an invitation", http.StatusBadRequest)
		return
	}
	isMember, err = database.GroupMember(invite.ReceiverID, invite.GroupID)
	if err != nil {
		helpers.HTTPError(w, "error checking if user is a member", http.StatusBadRequest)
		return
	}
	if isMember {
		helpers.HTTPError(w, "user already a member", http.StatusBadRequest)
		return
	}
	inviteID, err := database.CreateInvite(invite.GroupID, userID, invite.ReceiverID)
	if err != nil {
		helpers.HTTPError(w, "Failed to create invitation", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	log.Print(inviteID)
	// database.AddToNotificationTable(inviteID)
}

/*
InvitationResponseHandler responds to an invitation to a certain group.
This function responds to an invitation using the status
provided in the request body.
It requires a valid user session to respond to an invite.

Example:
	 // To respond to an invite
	 POST /api/invitationresponse
		Body:{
			"group_id":0
			"response":"accepted" | "rejected"
			}
*/
func InvitationResponseHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	var response models.GroupResponse
	err := json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		helpers.HTTPError(w, err.Error(), http.StatusBadRequest)
		return
	}
	if response.Status != "accepted" && response.Status != "rejected" {
		helpers.HTTPError(w, "status can only be rejected or accepted", http.StatusBadRequest)
		return
	}
	groupExists := database.CheckGroupID(response.GroupID)
	if !groupExists {
		helpers.HTTPError(w, "group ID does not exist", http.StatusBadRequest)
		return
	}

	isMember, err := database.GroupMember(userID, response.GroupID)
	if err != nil {
		helpers.HTTPError(w, "error checking if user is a member", http.StatusBadRequest)
		return
	}
	if isMember {
		helpers.HTTPError(w, "user already a member", http.StatusBadRequest)
		return
	}
	err = database.RespondToInvite(response, userID)
	if err != nil {
		helpers.HTTPError(w, "Failed to respond to invite", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	// database.UpdateNotificationTable()
}
