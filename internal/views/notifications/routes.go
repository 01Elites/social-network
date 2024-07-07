package notifications

import (
	"net/http"
	"social-network/internal/views/middleware"
)

func SetupFollowRoutes() {
	http.HandleFunc("POST /api/followrequest", middleware.AllowCORS(middleware.CheckAuth(FollowHandler)))
	http.HandleFunc("POST /api/followresponse", middleware.AllowCORS(middleware.CheckAuth(RespondToFollowHandler)))
}