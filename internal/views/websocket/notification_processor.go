package websocket

import (
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
			if err := sendMessageToWebSocket(clients[FollowRequest.ToUser].Conn, "NOTIFICATION", FollowRequest); err != nil {
				log.Println("Error sending follow request notification to WebSocket:", err)
				return
			}
		case GroupInvite := <-GroupInviteChan:
			if err := sendMessageToWebSocket(user.Conn, "NOTIFICATION", GroupInvite); err != nil {
				log.Println("Error sending TYPING to WebSocket:", err)
				return
			}
		case JoinRequest := <-JoinRequestChan:
			if err := sendMessageToWebSocket(clients[JoinRequest.ToUser].Conn, "NOTIFICATION", JoinRequest); err != nil {
				log.Println("Error sending TYPING to WebSocket:", err)
				return
			}
		case Event := <-EventChan:
			if err := sendMessageToWebSocket(user.Conn, "NOTIFICATION", Event); err != nil {
				log.Println("Error sending TYPING to WebSocket:", err)
				return
			}
		}
	}
}

func SendNotificationToChannel(notification types.Notification, notificationChan chan types.Notification) {
	if len(clients) == 0 {
		return
	}
	if clients[notification.ToUser] == nil {
		log.Println("User not online")
		return
	}
	notificationChan <- notification
}

func SendGroupRequestNotification(notification types.Notification) {
	if len(clients) == 0 {
		return
	}
	if clients[notification.ToUser] == nil {
		log.Println("User not online")
		return
	}
	JoinRequestChan <- notification
}

func SendGroupEventNotification(notification types.Notification) {
	if len(clients) == 0 {
		return
	}
	if clients[notification.ToUser] == nil {
		log.Println("User not online")
		return
	}
	EventChan <- notification
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
		case "REQUEST_TO_JOIN_GROUP":
			SendGroupRequestNotification(notification)
		case "EVENT":
			SendGroupEventNotification(notification)
		}
	}
	return nil
}
