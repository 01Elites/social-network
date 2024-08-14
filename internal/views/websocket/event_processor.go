package websocket

import (
	"encoding/json"
	"log"

	"social-network/internal/views/websocket/types"
	"social-network/internal/views/websocket/types/event"

	"github.com/gorilla/websocket"
)

func ProcessEvents(user *types.User) {
	// Send all the notifications in the database to the user
	go ProcessNotifications(user)
	// if err := SendUsersNotifications(user.ID); err != nil {
	// 	log.Printf("Error sending notifications: %v", err)
	// 	return
	// }
	defer func() {
		SetClientOffline(user)
	}()
	for {
		// Read message from WebSocket connection
		messageType, rawMessage, err := user.Conn.ReadMessage()
		// fmt.Println(user)
		if err != nil {
			log.Println("Error reading message from WebSocket:", err)
			// fmt.Println("user after", user)
			return
		}

		// if the messageType is not string will skip this message
		if messageType != websocket.TextMessage {
			continue
		}

		// Deserialize the message into the Event struct
		var message types.Event
		err = json.Unmarshal(rawMessage, &message)
		if err != nil {
			log.Println("Error unmarshalling JSON message into Event struct:", err)
			return
		}

		// Handle the event based on its type
		switch message.Type {

		case event.CHAT_OPENED:
			OpenChat(message, user)
		case event.CHAT_CLOSED:
			CloseChat(user)
		case event.SEND_MESSAGE:
			// Call function for NOTIFICATION
			if user.ChatOpened == "" {
				return
			} else if user.ChatOpenedIsGroup {
				SendMessageToGroup(message, user)
			} else {
				SendMessage(message, user)
			}
		case event.TYPING:
			// Call function for TYPING
			if user.ChatOpened == "" {
				return
			} else if user.ChatOpenedIsGroup {
				Typing(message, user, true)
			} else {
				Typing(message, user, false)
			}
		case event.USERLIST:
			GetUserList(user)
		case event.GET_NOTIFICATIONS:
			// Send all the notifications in the database to the user
			if err := SendUsersNotifications(user.ID); err != nil {
				log.Printf("Error sending notifications: %v", err)
			}
		// case event.GET_MESSAGES:
		// 	// GetMessages(message, user)
		case event.NOTIFICATION_READ:
			Notification(message, user)
		default:
			log.Println("Unknown event type:", message.Type)
		}
	}
}

/// todo:
// 1. Implement the Typing function
// 2. Implement the OpenChat function
// 3. Implement the CloseChat function
// 4. Implement the GetMessages function
// 5. Implement the GetNotifications function
// 6. get the list of users {online, offline,Dm's } and send it to the client
// 7. get the list of groups {online, offline, messages} and send it to the client

// to do with the notifications
// 1. Implement the notification function {send notification to the user by websocket}
// 2. notifications include:
// 	- new comment on the my post {post_id, commenter , comment}
// 	- new like on the my post {post_id, liker}
// 	- new friend request "follow request" {follower} to be [accepted or rejected]
// 	- new message in DM {sender, message}
// 	- new message in the group chat {sender, message , group_id}
//  - invitation to the group
//  - new reuest to join the group {which the user is the admin}
//  - new event in the group {which the user is the admin} MAYBE !!!

// required notification types:
// follow request
// group invitation
// request to join group
// event in the group
