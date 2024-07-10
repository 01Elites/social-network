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

func SetClientOnline(conn *websocket.Conn, user *types.User) {
	// Add the client to the Clients map
	cmutex.Lock()
	clients[user.Username] = user
	cmutex.Unlock()
	followees, err := database.GetUsersFollowees(user.UserID)
	if err != nil {
		log.Print(err)
	} else {
		for followee := range followees {
			sendMessageToWebSocket(conn, "USERLIST", data)
		}
	}
}