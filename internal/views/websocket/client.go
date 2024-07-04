package websocket

import "github.com/gorilla/websocket"

func SetClientOffline(userID string) {
	// Remove the client from the Clients map
	cmutex.Lock()
	delete(clients, userID)
	cmutex.Unlock()
}

func SetClientOnline(conn *websocket.Conn, userID string) {
	// Add the client to the Clients map
	cmutex.Lock()
	clients[userID] = conn
	cmutex.Unlock()
}
