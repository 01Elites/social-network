package websocket

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"

	"social-network/internal/helpers"
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
	token := r.URL.Query().Get("token")
	// Validate the session token
	userID, err := database.ValidateSessionToken(token)
	if err != nil {
		helpers.HTTPError(w, "Invalid session token", http.StatusUnauthorized)
		return
	}

	// Upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		helpers.HTTPError(w, "Could not upgrade to WebSocket", http.StatusInternalServerError)
		return
	}

	user, err := SetClientOnline(userID, conn)
	if err != nil {
		log.Println("Error setting client online:", err)
		helpers.HTTPError(w, "Could not set client online", http.StatusInternalServerError)
		return
	}
	updateUserInUserList(user, types.State.Online)

	go ProcessEvents(conn, user)
}

// Function to send JSON data to a WebSocket connection
func sendMessageToWebSocket(user *types.User, eventType string, data interface{}) error {
	// Check if the WebSocket connection is nil
	if user != nil && len(user.Conns) == 0 {
		log.Println("NO Connection")
		return nil
	}
	// Format the message payload according to the given event type and data
	eventMessage := types.Event{
		Type:    eventType,
		Payload: data,
	}
	// fmt.Println(eventMessage)
	messagesJSON, err := json.Marshal(eventMessage)
	if err != nil {
		log.Println("Error marshalling messages to JSON:", err)
		return nil
	}

	// Use the mutex to ensure only one goroutine writes at a time
	user.Mutex.Lock()
	defer user.Mutex.Unlock()
	// Write the payload to the WebSocket connection as a text message
	// Iterate over all connections and send the message
	for conn := range user.Conns {
		if err := conn.WriteMessage(websocket.TextMessage, messagesJSON); err != nil {
			log.Printf("Error writing message to WebSocket connection: %v\n", err)
		}
	}
	return nil
}
