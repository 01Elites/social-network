package events

import (
	"net/http"
	"social-network/internal/views/middleware"
)

func SetupEventsRoutes() {
	http.HandleFunc("GET /api/myevents", middleware.AllowCORS(middleware.CheckAuth((MyEventsHandler))))
}