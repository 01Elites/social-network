package websocket

import (
	"encoding/json"
	"fmt"
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
	// get the token from the header wedsocket protocol
	protocols := websocket.Subprotocols(r)
	if len(protocols) == 0 {
		helpers.HTTPError(w, "Authorization token required", http.StatusUnauthorized)
		return
	}
	token := protocols[0]
	userID, err := database.ValidateSessionToken(token)
	if err != nil {
		helpers.HTTPError(w, "Invalid session token", http.StatusUnauthorized)
		return
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Error upgrading to WebSocket:", err)
		helpers.HTTPError(w, "Could not upgrade to WebSocket", http.StatusInternalServerError)
		return
	}
	username, err := database.GetUserNameByID(userID)
	if err != nil {
		log.Println("Error getting user name:", err)
		helpers.HTTPError(w, "Could not get user name", http.StatusInternalServerError)
		return
	}
	user := types.User{
		ID:       userID,
		Username: username,
		Conn:     conn,
		Mutex:    &sync.Mutex{},
	}
	SetClientOnline(&user)
	go ProcessNotifications(&user)
	go ProcessEvents(&user)

	// send all the notifications in database to the user
	err = SendUsersNotifications(user.ID)
	if err != nil {
		log.Printf("error sending notifications:%v", err)
	}
}

// Function to send JSON data to a WebSocket connection
func sendMessageToWebSocket(user *types.User, eventType string, data interface{}) error {
	// Check if the WebSocket connection is nil
	if user.Conn == nil {
		log.Println("Connection is nil")
		return nil
	}
	fmt.Println(eventType, data)
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

	// Use the mutex to ensure only one goroutine writes at a time
	user.Mutex.Lock()
	defer user.Mutex.Unlock()
	// Write the payload to the WebSocket connection as a text message
	if err := user.Conn.WriteMessage(websocket.TextMessage, []byte(messagesJSON)); err != nil {
		log.Println("Error writing message to WebSocket:", err)
		return err
	}
	return nil
}
