package websocket

import (
	"log"

	"social-network/internal/views/websocket/types"

	"github.com/gorilla/websocket"
)

// Global channels for notifications
var (
	FollowRequestChan = make(chan types.Event)
	GroupInviteChan   = make(chan types.Event)
	JoinRequestChan   = make(chan types.Event)
	EventChan         = make(chan types.Event)
)

func ProcessNotifications(conn *websocket.Conn, username string) {
	// Get the notifications for the user
	// notifications, err := database.GetUserNotifications(username)
	// if err != nil {
	// 	log.Println("Error getting user notifications:", err)
	// 	return
	// }

	// // Send the notifications to the client
	// sendMessageToWebSocket(conn, event.NOTIFICATION, notifications)
	for {
		select {
		case FollowRequest := <-FollowRequestChan:
			if err := sendMessageToWebSocket(conn, "notificationType1", FollowRequest.Payload); err != nil {
				log.Println("Error sending SEND_MESSAGE to WebSocket:", err)
			}
		case GroupInvite := <-GroupInviteChan:
			if err := sendMessageToWebSocket(conn, "notificationType2", GroupInvite.Payload); err != nil {
				log.Println("Error sending TYPING to WebSocket:", err)
			}
		case JoinRequest := <-JoinRequestChan:
			if err := sendMessageToWebSocket(conn, "notificationType3", JoinRequest.Payload); err != nil {
				log.Println("Error sending TYPING to WebSocket:", err)
			}
		case Event := <-EventChan:
			if err := sendMessageToWebSocket(conn, "notificationType4", Event.Payload); err != nil {
				log.Println("Error sending TYPING to WebSocket:", err)
			}
		}
	}
}
