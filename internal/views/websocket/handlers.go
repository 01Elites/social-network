package websocket

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"social-network/internal/helpers"
	"social-network/internal/views/middleware"

	database "social-network/internal/database/querys"

	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	clients = make(map[string]*websocket.Conn)
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
