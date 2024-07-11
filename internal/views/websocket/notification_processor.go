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
			if err := sendMessageToWebSocket(user.Conn, "notificationType2", GroupInvite.Payload); err != nil {
				log.Println("Error sending TYPING to WebSocket:", err)
			}
		case JoinRequest := <-JoinRequestChan:
			if err := sendMessageToWebSocket(user.Conn, "notificationType3", JoinRequest.Payload); err != nil {
				log.Println("Error sending TYPING to WebSocket:", err)
			}
		case Event := <-EventChan:
			if err := sendMessageToWebSocket(user.Conn, "notificationType4", Event.Payload); err != nil {
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
