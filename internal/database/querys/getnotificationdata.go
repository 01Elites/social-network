package querys

import (
	"errors"
	"log"
	"social-network/internal/models"
	"social-network/internal/views/websocket/types"
	"time"
)

func GetFollowRequestNotification(request models.Request) (*types.Notification, error) {
	sender, err := GetUserProfile(request.Sender)
	if err != nil {
		log.Println("Failed to get user profile")
		return nil, err
	}
	recieverUsername, err := GetUserNameByID(request.Receiver)
	if err != nil {
		log.Println("Failed to get username of reciever")
		return nil, err
	}
	if request.CreatedAt.IsZero() {
		request.CreatedAt = time.Now()
	}

	notification := OrganizeFollowRequest(recieverUsername, *sender, request.CreatedAt)
	return &notification, nil
}

func GetGroupRequestData(userID string, requestID int) (*types.Notification, error) {
	groupID, groupTitle, groupCreatorID, requesterID, createdAt, err := getGroupFromRequest(requestID)
	if err != nil {
		if err == errors.New("no rows in result set") {
			log.Print("error getting groupID")
		}
		return nil, err
	}
	groupCreator, err := GetUserNameByID(groupCreatorID)
	if err != nil {
		log.Print("error getting group creator username")
		return nil, err
	}
	user, err := GetUserProfile(requesterID)
	if err != nil {
		log.Print("error getting user profile")
		return nil, err
	}
	notification := OrganizeGroupRequest(groupCreator, groupTitle, groupID, *user, createdAt)
	return &notification, nil
}

func GetGroupEventData(userID string, eventID int) (*types.Notification, error) {
	username, err := GetUserNameByID(userID)
	if err != nil {
		log.Print("failed to get username", err)
		return nil, err
	}
	options, err := GetEventOptions(eventID)
	if err != nil {
		log.Print("error getting event options")
		return nil, err
	}
	title, description, groupID, eventTime, firstname, lastname, err := GetEventDetails(eventID)
	if err != nil {
		log.Print("error getting event title", err)
		return nil, err
	}
	groupTitle, err := GetGroupTitle(groupID)
	if err != nil {
		log.Print("error getting group title", err)
		return nil, err
	}
	ChoiceMade, err := MadeChoice(eventID, userID)
	if ChoiceMade {
		return nil, errors.New("user already made a choice")
	}
	if err != nil {
		log.Print("error getting event choices", err)
		return nil, err
	}

	eventDetails := types.EventDetails{
		ID:          eventID,
		Title:       title,
		Options:     options,
		EventTime:   eventTime,
		Description: description,
		Creator: firstname + " " + lastname,
	}
	notification := OrganizeGroupEventRequest(username, groupTitle, groupID, eventDetails)
	return &notification, nil
}

func GetGroupInvitationData(userID string, invitationID int) (*types.Notification, error) {
	invitedUserID, groupID, groupTitle, inviter, err := getGroupFromInvitation(invitationID)
	if err != nil {
		if err == errors.New("no rows in result set") {
			log.Print("error getting groupID")
		}
		return nil, err
	}
	ifMember, err := GroupMember(invitedUserID, groupID)
	if err != nil {
		log.Print("error checking if user is a member")
		return nil, err
	}
	if ifMember {
		return nil, errors.New("user is already a member")
	}

	invitedUser, err := GetUserNameByID(invitedUserID)
	if err != nil {
		log.Print("error getting user name")
		return nil, err
	}

	notification := OrganizeGroupInvitation(invitedUser, groupID, groupTitle, inviter)
	return &notification, nil
}
