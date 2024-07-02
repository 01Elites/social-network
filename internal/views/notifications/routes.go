package notifications

import (
	"net/http"
	"social-network/internal/views/middleware"
)

func SetupNotificationRoutes() {
	http.HandleFunc("POST /api/followrequest/{id}", middleware.ValidateSessionMiddleware(CreateFollowHandler))
	http.HandleFunc("POST /api/followresponse/{id}", middleware.ValidateSessionMiddleware(RespondToFollowHandler))
}