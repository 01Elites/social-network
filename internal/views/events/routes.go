package events

import (
	"net/http"
	"social-network/internal/views/middleware"
)

func SetupFriendsRoutes() {
	http.HandleFunc("GET /api/myevents", middleware.AllowCORS(middleware.CheckAuth((MyEventsHandler))))
}