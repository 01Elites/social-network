package notifications

import (
	"net/http"
	"social-network/internal/views/middleware"
)

func SetupNotificationRoutes() {
	http.HandleFunc("POST /api/followrequest", middleware.AllowCORS(middleware.CheckAuth(CreateFollowHandler)))
	http.HandleFunc("POST /api/followresponse", middleware.AllowCORS(middleware.CheckAuth(RespondToFollowHandler)))
}