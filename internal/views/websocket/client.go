package websocket

import (
	"log"

	database "social-network/internal/database/querys"
	"social-network/internal/views/websocket/types"

	"github.com/gorilla/websocket"
)

func SetClientOffline(userName string) {
	// Remove the client from the Clients map
	cmutex.Lock()
	delete(clients, userName)
	cmutex.Unlock()
}

func SetClientOnline(conn *websocket.Conn, userName string) {
	// Add the client to the Clients map
	cmutex.Lock()
	clients[userName] = &types.User{
		Username: userName,
		State:    "online",
		Conn:     conn,
	}
	cmutex.Unlock()
	followees, err := database.GetUsersFollowees(userName)
	if err != nil {
		log.Print(err)
	} else {
		for followee := range followees {
			sendMessageToWebSocket(conn, "USERLIST", data)
		}
	}
}
