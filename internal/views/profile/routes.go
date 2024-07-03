package profile

import (
	"net/http"
	"social-network/internal/views/middleware"
)

func SetupProfileRoutes() {
	http.HandleFunc("GET /api/profile", middleware.AllowCORS(middleware.CheckAuth(getProfile)))
	http.HandleFunc("PATCH /api/profile", middleware.AllowCORS(middleware.CheckAuth(patchProfile)))
	http.HandleFunc("GET /api/profile", middleware.ValidateSessionMiddleware(getProfile))
	http.HandleFunc("PATCH /api/profile", middleware.ValidateSessionMiddleware(patchProfile))
	http.HandleFunc("GET /api/profile/{id}", middleware.ValidateSessionMiddleware(getProfileByID))
	http.HandleFunc("GET /api/profile/{id}/posts", middleware.ValidateSessionMiddleware(getProfilePosts))
}