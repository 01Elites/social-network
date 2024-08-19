package websocket

import (
	"encoding/json"
	"fmt"
	"log"

	"social-network/internal/views/websocket/types"
	"social-network/internal/views/websocket/types/event"

	"github.com/gorilla/websocket"
)

func ProcessEvents(conn *websocket.Conn,user *types.User) {
	// Send all the notifications in the database to the user
	defer func() {
		SetClientOffline(user, conn)
		updateUserInUserList(user, types.State.Offline)
	}()
	for {
		// Read message from WebSocket connection
		messageType, rawMessage, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				return
			}
			log.Println("Error reading message from WebSocket:", err)
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
			go ProcessNotifications(user)
			if err := SendUsersNotifications(user.ID); err != nil {
				log.Printf("Error sending notifications: %v", err)
			}
		case event.NOTIFICATION_READ:
			Notification(message, user)
		default:
			log.Println("Unknown event type:", message.Type)
		}
	}
}
