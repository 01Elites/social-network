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
		default:
			log.Println("Unknown event type:", message.Event)
		}
	}
}
