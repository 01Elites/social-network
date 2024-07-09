package websocket

import "github.com/gorilla/websocket"

func SetClientOffline(userName string) {
	// Remove the client from the Clients map
	cmutex.Lock()
	delete(clients, userName)
	cmutex.Unlock()
}

func SetClientOnline(conn *websocket.Conn, userName string) {
	// Add the client to the Clients map
	cmutex.Lock()
	clients[userName] = conn
	cmutex.Unlock()
}
