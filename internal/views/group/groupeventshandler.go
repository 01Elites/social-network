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

	if event.GroupID == 0 || event.Title == "" || event.EventTime.IsZero() {
		http.Error(w, "missing required fields", http.StatusBadRequest)
		return
	}
	groupExists := database.CheckGroupID(event.GroupID)
	if !groupExists {
		helpers.HTTPError(w, "group ID does not exist", http.StatusBadRequest)
		return
	}
	isMember, err := database.GroupMember(userID, event.GroupID)
	if err != nil {
		helpers.HTTPError(w, "error checking if user is a member", http.StatusBadRequest)
		return
	}
	if !isMember {
		helpers.HTTPError(w, "you are not a memeber", http.StatusBadRequest)
		return
	}

	var eventID int
	eventID, err = database.CreateEvent(event.GroupID, userID, event.Title, event.Description, event.EventTime)
	if err != nil {
		helpers.HTTPError(w, "failed to create event", http.StatusNotFound)
		return
	}

	err = database.CreateEventOptions(eventID, event.Options)
	if err != nil {
		helpers.HTTPError(w, "failed to create event options", http.StatusNotFound)
		return
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
	w.WriteHeader(http.StatusOK)
}

/*
func CancelEventHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	var event models.ID
	err := json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		helpers.HTTPError(w, "failed to decode request", http.StatusBadRequest)
		return
	}
	groupId := database.CheckEventID(event.ID)
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
	err = database.CancelEvent(event.ID)
	if err != nil {
		helpers.HTTPError(w, "error when canceling event", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

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
