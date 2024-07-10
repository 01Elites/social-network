package websocket

import (
	"log"

	"social-network/internal/views/websocket/types"

	"github.com/gorilla/websocket"
)

// Global channels for notifications
var (
	Channel1 = make(chan types.Event)
	Channel2 = make(chan types.Event)
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
		case notification1 := <-Channel1:
			if err := sendMessageToWebSocket(conn, "notificationType1", notification1.Payload); err != nil {
				log.Println("Error sending SEND_MESSAGE to WebSocket:", err)
			}
		case notification2 := <-Channel2:
			if err := sendMessageToWebSocket(conn, "notificationType2", notification2.Payload); err != nil {
				log.Println("Error sending TYPING to WebSocket:", err)
			}
		}
	}
}
