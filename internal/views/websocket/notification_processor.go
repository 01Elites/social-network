package websocket

import (
	"log"

	"social-network/internal/models"
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
	// notifications, err := database.GetUserNotifications()
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
			if err := sendMessageToWebSocket(user.Conn, "NOTIFICATION", GroupInvite); err != nil {
				log.Println("Error sending TYPING to WebSocket:", err)
			}
		case JoinRequest := <-JoinRequestChan:
			conn := clients[JoinRequest.ToUser].Conn
			log.Print(JoinRequest)
			if conn == nil {
				log.Println("User not online")
				return
			}
			if err := sendMessageToWebSocket(conn, "NOTIFICATION", JoinRequest); err != nil {
				log.Println("Error sending TYPING to WebSocket:", err)
			}
		case Event := <-EventChan:
			if err := sendMessageToWebSocket(user.Conn, "NOTIFICATION", Event); err != nil {
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

func GroupRequestNotification(groupCreator string, GroupTitle string, groupID int, requester models.UserProfile) {
	notification := types.Notification{
		Type:    "REQUEST_TO_JOIN_GROUP",
		Message: "You have a new group request",
		ToUser:  groupCreator,
		Metadata: types.GroupRequestMetadata{
			UserDetails: types.UserDetails{
				Username:  requester.Username,
				FirstName: requester.FirstName,
				LastName:  requester.LastName,
			},
			Group: types.GroupNotification{
				ID:    groupID,
				Title: GroupTitle,
			},
		},
	}
	err := database.AddToNotificationTable(groupCreator, "join_request", groupID)
	if err !=nil {
		log.Println("error adding notification to database")
				return
	}
	JoinRequestChan <- notification
}
