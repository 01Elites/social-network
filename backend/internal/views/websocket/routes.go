package websocket

import (
	"net/http"

	"social-network/internal/views/middleware"
)

func SetupWebSocketRoutes() {
	http.HandleFunc("/api/ws", middleware.AllowCORS((HandleWebSocket)))
}
