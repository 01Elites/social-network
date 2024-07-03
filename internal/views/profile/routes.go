package profile

import (
	"net/http"
	"social-network/internal/views/middleware"
)

func SetupProfileRoutes() {
	http.HandleFunc("GET /api/profile", middleware.AllowCORS(middleware.CheckAuth(getProfile)))
	http.HandleFunc("PATCH /api/profile", middleware.AllowCORS(middleware.CheckAuth(patchProfile)))
}
