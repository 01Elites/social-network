package profile

import (
	"net/http"
	"social-network/internal/views/middleware"
)

func SetupProfileRoutes() {
	http.HandleFunc("GET /api/profile", middleware.ValidateSessionMiddleware(getProfile))
	http.HandleFunc("PATCH /api/profile", middleware.ValidateSessionMiddleware(patchProfile))
}