package notifications

import (
	"net/http"
	"social-network/internal/views/middleware"
)

func SetupNotificationRoutes() {
	http.HandleFunc("POST /api/followrequest", middleware.ValidateSessionMiddleware(CreateFollowHandler))
	http.HandleFunc("POST /api/followresponse", middleware.ValidateSessionMiddleware(RespondToFollowHandler))
}