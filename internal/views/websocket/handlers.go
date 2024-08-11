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

	// Retrieve the username from the database
	username, err := database.GetUserNameByID(userID)
	if err != nil {
		log.Println("Error getting user name:", err)
		helpers.HTTPError(w, "Could not get user name", http.StatusInternalServerError)
		return
	}

	// Create the user object and mark the user as online
	user := types.User{
		ID:       userID,
		Username: username,
		Conn:     conn,
		Mutex:    &sync.Mutex{},
	}
	SetClientOnline(&user)
	defer SetClientOffline(user.Username) // Ensure the user is marked offline on function exit

	// Use a WaitGroup to synchronize goroutines
	var wg sync.WaitGroup
	wg.Add(2)

	// Goroutine to process notifications
	go func() {
		defer wg.Done()
		ProcessNotifications(&user)
	}()

	// Goroutine to process events
	go func() {
		defer wg.Done()
		ProcessEvents(&user)
	}()

	// Send all the notifications in the database to the user
	if err := SendUsersNotifications(user.ID); err != nil {
		log.Printf("Error sending notifications: %v", err)
	}

	// Wait for both goroutines to finish before closing the WebSocket connection
	wg.Wait()
}
// Function to send JSON data to a WebSocket connection
func sendMessageToWebSocket(user *types.User, eventType string, data interface{}) error {
	// Check if the WebSocket connection is nil
	if user.Conn == nil {
		log.Println("Connection is nil")
		return nil
	}
	// Format the message payload according to the given event type and data
	eventMessage := types.Event{
		Type:    eventType,
		Payload: data,
	}
	fmt.Println(eventMessage)
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
