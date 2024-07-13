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
			if err := sendMessageToWebSocket(user.Conn, "notificationType1", FollowRequest); err != nil {
				log.Println("Error sending SEND_MESSAGE to WebSocket:", err)
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

func SendUsersNotifications(userID string)error{
	notifications, err := database.GetUserNotifications(userID)
	if err != nil {
		log.Println("Error getting user notifications:", err)
		return err
	}
	for _, notification := range notifications {
	switch notification.Type{
	case "FOLLOW_REQUEST":
		
	case "GROUP_INVITATION":
	case "REQUEST_TO_JOIN_GROUP":
		SendGroupRequestNotification(notification)
	case "EVENT":
		SendGroupEventNotification(notification)
	}
}
return nil
}