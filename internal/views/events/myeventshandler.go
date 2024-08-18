package events

import (
	"encoding/json"
	"net/http"
	"time"

	database "social-network/internal/database/querys"
	"social-network/internal/helpers"
	"social-network/internal/models"
	"social-network/internal/views/middleware"
)

func MyEventsHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	username, err := database.GetUserNameByID(userID)
	if err != nil {
		helpers.HTTPError(w, "failed to get username", http.StatusNotFound)
		return
	}
	groupIDs, err := database.GetAllGroupIDs()
	if err != nil {
		helpers.HTTPError(w, "failed to get group ids", http.StatusNotFound)
		return
	}
	var myEvents models.MyEvents
		for _, groupID := range groupIDs {
			isMember, err := database.GroupMember(userID, groupID)
			if err != nil {
				helpers.HTTPError(w, "error checking if user is a member", http.StatusBadRequest)
				return
			}
			if !isMember {
				continue
			}
			events, err := database.GetGroupEvents(groupID)
			if err != nil {
				helpers.HTTPError(w, "Failed to get group events", http.StatusInternalServerError)
				return
			}
			for _, event := range events {
				if event.EventTime.Before(time.Now()){
					myEvents.PastEvents = append(myEvents.PastEvents, event)
					continue
				}
				if len(event.Options) < 2 {
					continue
				}
				var optionChoosen bool
				for _, eventUsername := range event.Options[0].Usernames {
					if username == eventUsername {
						optionChoosen = true
						myEvents.UpcomingEvents = append(myEvents.UpcomingEvents, event)
						break	
					}
				}
				if optionChoosen{
					continue
				} else {
					for _, eventUsername := range event.Options[1].Usernames {
						if username == eventUsername {
							optionChoosen = true
							myEvents.UpcomingEvents = append(myEvents.UpcomingEvents, event)
							break	
						}
					}
				}
				myEvents.PendingEvents = append(myEvents.PendingEvents, event)
				}
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(myEvents)
}
