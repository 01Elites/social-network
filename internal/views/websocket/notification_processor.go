package websocket

import (
	"encoding/json"
	"log"

	database "social-network/internal/database/querys"
	"social-network/internal/views/websocket/types"
)

// Global channels for notifications
var (
	FollowRequestChan = make(chan types.Notification)
	GroupInviteChan   = make(chan types.Notification)
	JoinRequestChan   = make(chan types.Notification)
	EventChan         = make(chan types.Notification)
)

func ProcessNotifications(user *types.User) {
	// Get the notifications for the user
	// // Send the notifications to the client
	// sendMessageToWebSocket(conn, event.NOTIFICATION, notifications)
	for {
		select {
		case FollowRequest := <-FollowRequestChan:
			if err := sendMessageToWebSocket(Clients[FollowRequest.ToUser], "NOTIFICATION", FollowRequest); err != nil {
				log.Println("Error sending follow request notification to WebSocket:", err)
				return
			}
		case GroupInvite := <-GroupInviteChan:
			if err := sendMessageToWebSocket(Clients[GroupInvite.ToUser], "NOTIFICATION", GroupInvite); err != nil {
				log.Println("Error sending Group Invite to WebSocket:", err)
				return
			}
		case JoinRequest := <-JoinRequestChan:
			if err := sendMessageToWebSocket(Clients[JoinRequest.ToUser], "NOTIFICATION", JoinRequest); err != nil {
				log.Println("Error sending Join Request to WebSocket:", err)
				return
			}
		case Event := <-EventChan:
			if err := sendMessageToWebSocket(Clients[Event.ToUser], "NOTIFICATION", Event); err != nil {
				log.Println("Error sending Event to WebSocket:", err)
				return
			}
		}
	}
}

func SendNotificationToChannel(notification types.Notification, notificationChan chan types.Notification) {
	if len(Clients) == 0 {
		return
	}
	if Clients[notification.ToUser] == nil {
		log.Println("User not online", notification.ToUser)
		return
	}
	notificationChan <- notification
}

func SendUsersNotifications(userID string) error {
	notifications, err := database.GetUserNotifications(userID)
	if err != nil {
		log.Println("Error getting user notifications:", err)
		return err
	}
	for _, notification := range notifications {
		switch notification.Type {
		case "FOLLOW_REQUEST":
			SendNotificationToChannel(notification, FollowRequestChan)
		case "GROUP_INVITATION":
			SendNotificationToChannel(notification, GroupInviteChan)
		case "REQUEST_TO_JOIN_GROUP":
			SendNotificationToChannel(notification, JoinRequestChan)
		case "EVENT":
			SendNotificationToChannel(notification, EventChan)
		}
	}
	return nil
}

// Notification function to handle notification events
func Notification(RevEvent types.Event, user *types.User) {
	// Convert map to JSON
	jsonPayload, err := json.Marshal(RevEvent.Payload)
	if err != nil {
		log.Println("Error marshaling payload to JSON:", err)
		return
	}

	// Unmarshal event payload to get recipient
	var payload struct {
		NotificationID int `json:"notification_id"`
	}

	if err := json.Unmarshal(jsonPayload, &payload); err != nil {
		log.Println("Error decoding event typing:", err)
		return
	}

	// Set notification as read in database
	if err := database.SetNotificationAsRead(payload.NotificationID); err != nil {
		log.Println("Error setting notification as read:", err)
		return
	}

}
