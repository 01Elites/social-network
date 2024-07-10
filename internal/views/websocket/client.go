package websocket

import ("github.com/gorilla/websocket"
				database "social-network/internal/database/querys"
				"log"
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
	clients[userName] = conn
	cmutex.Unlock()
	followees, err := database.GetUsersFollowees(userName)
	if err != nil {
		log.Print(err)
	}else{
	for followee, _ := range followees {
		sendMessageToWebSocket(conn, "USERLIST", data)
	}
	}
}
