package websocket

import (
	"encoding/json"
	"log"
	"social-network/internal/views/websocket/events_type"

	"github.com/gorilla/websocket"
)

func ProcessEvents(conn *websocket.Conn, userName string) {
	defer func() {
		// Remove the client from the Clients map when the connection is closed
		conn.Close()
		SetClientOffline(userName)
	}()

	for {
		// Read message from WebSocket connection
		messageType, rawMessage, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message from WebSocket:", err)
			return
		}

		// if the messageType is not string will skip this message
		if messageType != websocket.TextMessage {
			continue
		}

		// Deserialize the message into the Event struct
		var message events_type.Message
		err = json.Unmarshal(rawMessage, &message)
		if err != nil {
			log.Println("Error unmarshalling JSON message into Event struct:", err)
			return
		}

		// Handle the event based on its type
		switch message.Event {
		case events_type.SEND_MESSAGE:
			// Call function for NOTIFICATION
			log.Println("SEND_MESSAGE")
		case events_type.TYPING:
			// Call function for TYPING
			// Typing(event, user)
		case events_type.CHAT_OPENED:
			// OpenChat(event, user)
		case events_type.CHAT_CLOSED:
			// CloseChat(user)
		case events_type.GET_MESSAGES:
			// GetMessages(event, user)
		case events_type.GET_NOTIFICATIONS:
			// GetNotifications(event, user)
		default:
			log.Println("Unknown event type:", message.Event)
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
