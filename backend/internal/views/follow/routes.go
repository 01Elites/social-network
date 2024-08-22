package follow

import (
	"net/http"
	"social-network/internal/views/middleware"
)

func SetupFollowRoutes() {
	http.HandleFunc("POST /api/follow_request", middleware.AllowCORS(middleware.CheckAuth(FollowHandler)))
	http.HandleFunc("POST /api/follow_response", middleware.AllowCORS(middleware.CheckAuth(RespondToFollowHandler)))
}
