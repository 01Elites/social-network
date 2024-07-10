package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"social-network/internal/helpers"
	"social-network/internal/views/middleware"
	"social-network/internal/views/websocket/types"

	database "social-network/internal/database/querys"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	clients = make(map[string]*types.User)
	cmutex  sync.Mutex
)

func HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		helpers.HTTPError(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	fmt.Printf("User %s connected to WebSocket\n", userID)
	// Upgrade HTTP connection to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		http.Error(w, "Could not open WebSocket connection", http.StatusBadRequest)
		return
	}
	userName, err := database.GetUserNameByID(userID)
	if err != nil {
		log.Println("Error getting user name:", err)
		http.Error(w, "Could not get user name", http.StatusInternalServerError)
		return
	}
	SetClientOnline(conn, userName)

	go ProcessEvents(conn, userName)
}

// Function to send JSON data to a WebSocket connection
func sendMessageToWebSocket(conn *websocket.Conn, eventType string, data interface{}) error {
	// Check if the WebSocket connection is nil
	if conn == nil {
		log.Println("Connection is nil")
		return nil
	}
	// Format the message payload according to the given event type and data
	eventMessage := types.Event{
		Type:    eventType,
		Payload: data,
	}
	messagesJSON, err := json.Marshal(eventMessage)
	if err != nil {
		log.Println("Error marshalling messages to JSON:", err)
		return nil
	}
	// Write the payload to the WebSocket connection as a text message
	if err := conn.WriteMessage(websocket.TextMessage, []byte(messagesJSON)); err != nil {
		log.Println("Error writing message to WebSocket:", err)
		return err
	}
	return nil
}
