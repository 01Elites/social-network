package querys

import (
	"log"

	"social-network/internal/models"
	"social-network/internal/views/websocket/types"
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
	notification := OrganizeFollowRequest(recieverUsername, *sender)
	return &notification, nil
}

func GetGroupRequestData(userID string, requestID int) (string, string, int, models.UserProfile) {
	groupID, groupTitle, err := getGroupFromRequest(requestID)
	if err != nil {
		log.Print("error getting groupID")
		return "", "", 0, models.UserProfile{}
	}
	user, err := GetUserProfile(userID)
	if err != nil {
		log.Print("error getting groupID")
		return "", "", 0, models.UserProfile{}
	}
	return user.Username, groupTitle, groupID, *user
}

func GetGroupEventData(userID string, eventID int) (string, string, int, types.EventDetails) {
	username, err := GetUserNameByID(userID)
	if err != nil {
		log.Print("failed to get username", err)
		return "", "", 0, types.EventDetails{}
	}

	options, err := GetEventOptions(eventID)
	if err != nil {
		log.Print("error getting event options")
		return "", "", 0, types.EventDetails{}
	}
	title, groupID, err := GetEventDetails(eventID)
	if err != nil {
		log.Print("error getting event title", err)
		return "", "", 0, types.EventDetails{}
	}
	groupTitle, err := GetGroupTitle(groupID)
	if err != nil {
		log.Print("error getting group title", err)
		return "", "", 0, types.EventDetails{}
	}
	eventDetails := types.EventDetails{
		ID:      eventID,
		Title:   title,
		Options: options,
	}
	return username, groupTitle, groupID, eventDetails
}
