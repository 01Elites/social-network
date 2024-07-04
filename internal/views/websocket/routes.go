package websocket

import (
	"net/http"

	"social-network/internal/views/middleware"
)

func SetupWebSocketRoutes() {
	http.HandleFunc("/ws", middleware.AllowCORS(middleware.CheckAuth(HandleWebSocket)))
}
