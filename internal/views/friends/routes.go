package friends

import (
	"net/http"
	"social-network/internal/views/middleware"
)

func SetupFriendsRoutes() {
	http.HandleFunc("GET /api/friends/{user_name}", middleware.AllowCORS(middleware.CheckAuth((GetFriendsHandler))))
}
