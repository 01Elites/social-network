package websocket

import (
	"log"

	"social-network/internal/views/websocket/types"
)

// Global channels for notifications
var (
	FollowRequestChan = make(chan types.Notification)
	GroupInviteChan   = make(chan types.Event)
	JoinRequestChan   = make(chan types.Event)
	EventChan         = make(chan types.Event)
)

func ProcessNotifications(user *types.User) {
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
			if err := sendMessageToWebSocket(user.Conn, "notificationType1", FollowRequest); err != nil {
				log.Println("Error sending SEND_MESSAGE to WebSocket:", err)
			}
		case GroupInvite := <-GroupInviteChan:
			if err := sendMessageToWebSocket(user.Conn, "NOTIFICATION", GroupInvite.Payload); err != nil {
				log.Println("Error sending TYPING to WebSocket:", err)
			}
		case JoinRequest := <-JoinRequestChan:
			conn := clients[JoinRequest.ToUser].Conn
			log.Print(JoinRequest)
			if conn == nil {
				log.Println("User not online")
				return
			}
			if err := sendMessageToWebSocket(conn, "NOTIFICATION", JoinRequest.Payload); err != nil {
				log.Println("Error sending TYPING to WebSocket:", err)
			}
		case Event := <-EventChan:
			if err := sendMessageToWebSocket(user.Conn, "NOTIFICATION", Event.Payload); err != nil {
				log.Println("Error sending TYPING to WebSocket:", err)
			}
		}
	}
}

// func FollowRequestNotification(request models.Request) {
// 	notification := types.Notification{
// 		Type:    "FOILLOW_REQUEST",
// 		Message: "You have a new follow request",
// 		Metadata: types.FollowRequestMetadata{
// 			UserDetails: types.UserDetails{
// 				Username: request.Sender,
// 			},
// 		},
// 	}
// 	FollowRequestChan <- notification
// }
