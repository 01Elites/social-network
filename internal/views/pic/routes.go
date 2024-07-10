package pic

import (
	"net/http"

	"social-network/internal/views/middleware"
)

// SetupImageRoutes sets up the routes for the image handlers
func SetupImageRoutes() {
	http.HandleFunc("GET /api/image/{path}", middleware.AllowCORS(middleware.CheckAuth(servesImage)))
}
