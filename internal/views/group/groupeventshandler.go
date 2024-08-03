package group

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	database "social-network/internal/database/querys"
	"social-network/internal/helpers"
	"social-network/internal/models"
	"social-network/internal/views/middleware"
	"social-network/internal/views/websocket"
	"social-network/internal/views/websocket/types"
)

func CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	var event models.CreateEvent
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		http.Error(w, "failed to decode request: "+err.Error(), http.StatusBadRequest)
		log.Printf("failed to decode request: %v", err)
		return
	}
	if event.GroupID == 0 || event.Title == "" || event.EventTime.IsZero() || len(event.Options) < 2 {
		http.Error(w, "missing required fields", http.StatusBadRequest)
		return
	}

	if event.EventTime.Before(time.Now().Add(24 * time.Hour)) {
		http.Error(w, "event time cannot be in the past", http.StatusBadRequest)
		return
	}

	if len(event.Title) > 15 || len(event.Description) > 200 {
		http.Error(w, "title or description too long", http.StatusBadRequest)
		return
	}

	groupTitle, err := database.GetGroupTitle(event.GroupID)
	if err != nil {
		helpers.HTTPError(w, "group ID does not exist", http.StatusBadRequest)
		return
	}
	isMember, err := database.GroupMember(userID, event.GroupID)
	if err != nil {
		helpers.HTTPError(w, "error checking if user is a member", http.StatusBadRequest)
		return
	}
	if !isMember {
		helpers.HTTPError(w, "you are not a member", http.StatusBadRequest)
		return
	}
	groupMembers, groupMembersIDs, err := database.GetGroupMembers(event.GroupID)
	if err != nil {
		helpers.HTTPError(w, "failed to get group members", http.StatusNotFound)
		return
	}
	eventID, err := database.CreateEvent(event.GroupID, userID, event.Title, event.Description, event.EventTime)
	if err != nil {
		helpers.HTTPError(w, "failed to create event", http.StatusNotFound)
		return
	}
	err = database.CreateEventOptions(eventID, event.Options)
	if err != nil {
		helpers.HTTPError(w, "failed to create event options", http.StatusNotFound)
		return
	}
	groupEvent := types.EventDetails{
		ID:      eventID,
		Title:   event.Title,
		Options: event.Options,
	}
	for i, member := range groupMembers {
		notificationID, err := database.AddToNotificationTable(groupMembersIDs[i], "event_notification", eventID)
		if err != nil {
			log.Println("error adding notification to database")
			return
		}
		notification := database.OrganizeGroupEventRequest(member.UserName, groupTitle, event.GroupID, groupEvent)
		notification.ID = notificationID
		websocket.SendNotificationToChannel(notification, websocket.JoinRequestChan)
	}
	w.WriteHeader(http.StatusOK)
}

func EventResponseHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	var response models.EventResp
	err := json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		helpers.HTTPError(w, "failed to decode request", http.StatusBadRequest)
		return
	}
	groupId := database.CheckEventID(response.EventID)
	if groupId == 0 {
		helpers.HTTPError(w, "Event ID does not exist", http.StatusBadRequest)
		return
	}
	isMember, err := database.GroupMember(userID, groupId)
	if err != nil {
		helpers.HTTPError(w, "error checking if user is a member", http.StatusBadRequest)
		return
	}
	if !isMember {
		helpers.HTTPError(w, "you are not a memeber", http.StatusBadRequest)
		return
	}
	err = database.RespondToEvent(response, userID)
	if err != nil {
		helpers.HTTPError(w, "error when responding to request", http.StatusNotFound)
		return
	}
	database.UpdateNotificationTable(response.EventID, "accepted", "event_notification", userID)
	w.WriteHeader(http.StatusOK)
}

func CancelEventHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	var event models.CancelEvent
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		helpers.HTTPError(w, "failed to decode request", http.StatusBadRequest)
		return
	}
	groupId := database.CheckEventID(event.EventID)
	if groupId == 0 {
		helpers.HTTPError(w, "Event ID does not exist", http.StatusBadRequest)
		return
	}
	isEventCreator, err := database.EventCreator(userID, event.EventID)
	if err != nil {
		helpers.HTTPError(w, "error checking if user is the creator", http.StatusBadRequest)
		return
	}
	if !isEventCreator {
		helpers.HTTPError(w, "you are not the creator of this event", http.StatusBadRequest)
		return
	}
	isMember, err := database.GroupMember(userID, groupId)
	if err != nil {
		helpers.HTTPError(w, "error checking if user is a member", http.StatusBadRequest)
		return
	}
	if !isMember {
		helpers.HTTPError(w, "you are not a memeber", http.StatusBadRequest)
		return
	}
	err = database.CancelEvent(event.EventID)
	if err != nil {
		helpers.HTTPError(w, "error when canceling event", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

/*
func RespondToEventOptionHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	var response models.EventResp
	err := json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		helpers.HTTPError(w, "failed to decode request", http.StatusBadRequest)
		return
	}
	groupId := database.CheckEventID(response.EventID)
	if groupId == 0 {
		helpers.HTTPError(w, "Event ID does not exist", http.StatusBadRequest)
		return
	}
	isMember, err := database.GroupMember(userID, groupId)
	if err != nil {
		helpers.HTTPError(w, "error checking if user is a member", http.StatusBadRequest)
		return
	}
	if !isMember {
		helpers.HTTPError(w, "you are not a memeber", http.StatusBadRequest)
		return
	}
	err = database.RespondToEventOption(response, userID)
	if err != nil {
		helpers.HTTPError(w, "error when responding to request", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetEventOptionsHandler(w http.ResponseWriter, r *http.Request) {
	var event models.ID
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		helpers.HTTPError(w, "failed to decode request", http.StatusBadRequest)
		return
	}
	options, err := database.GetEventOptions(event.ID)
	if err != nil {
		helpers.HTTPError(w, "failed to get event options", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(options)
}

func GetEventResponsesHandler(w http.ResponseWriter, r *http.Request) {
	var event models.ID
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		helpers.HTTPError(w, "failed to decode request", http.StatusBadRequest)
		return
	}
	options, err := database.GetEventResponses(event.ID)
	if err != nil {
		helpers.HTTPError(w, "failed to get event responses", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(options)
}

func GetEventResponsesByOptionHandler(w http.ResponseWriter, r *http.Request) {
	var event models.ID
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		helpers.HTTPError(w, "failed to decode request", http.StatusBadRequest)
		return
	}
	options, err := database.GetEventResponsesByOption(event.ID)
	if err != nil {
		helpers.HTTPError(w, "failed to get event responses", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(options)
}

func GetEventResponsesByUserHandler(w http.ResponseWriter, r *http.Request) {
	var event models.ID
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		helpers.HTTPError(w, "failed to decode request", http.StatusBadRequest)
		return
	}
	options, err := database.GetEventResponsesByUser(event.ID)
	if err != nil {
		helpers.HTTPError(w, "failed to get event responses", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(options)
}

func GetEventResponsesByOptionAndUserHandler(w http.ResponseWriter, r *http.Request) {
	var event models.ID
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		helpers.HTTPError(w, "failed to decode request", http.StatusBadRequest)
		return
	}
	options, err := database.GetEventResponsesByOptionAndUser(event.ID)
	if err != nil {
		helpers.HTTPError(w, "failed to get event responses", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(options)
}
*/

func GetMyGroupsHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	var groups models.Groups

	// get the id of all the groups in database
	groupIDs, err := database.GetAllGroupIDs()
	if err != nil {
		helpers.HTTPError(w, "failed to get group ids", http.StatusNotFound)
		return
	}

	if len(groupIDs) != 0 {
		for _, groupID := range groupIDs {
			group, err := database.GetGroupFeedInfo(groupID, userID)
			if err != nil {
				helpers.HTTPError(w, "failed to get group info", http.StatusNotFound)
				return
			}
			if group.ID != 0 {
				if group.IsCreator {
					groups.Owned = append(groups.Owned, group)
				} else if group.IsMember {
					groups.Joined = append(groups.Joined, group)
				} else {
					groups.Explore = append(groups.Explore, group)
				}
			}
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(groups)

}
