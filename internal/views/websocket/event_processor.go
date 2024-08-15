package websocket

import (
	"encoding/json"
	"fmt"
	"log"

	"social-network/internal/views/websocket/types"
	"social-network/internal/views/websocket/types/event"

	"github.com/gorilla/websocket"
)

func ProcessEvents(user *types.User) {
	// Send all the notifications in the database to the user
	go ProcessNotifications(user)
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

		fmt.Printf("Received message: %v\n", message)
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
