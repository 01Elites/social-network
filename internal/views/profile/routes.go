package profile

import (
	"net/http"
	"social-network/internal/views/middleware"
)

func SetupProfileRoutes() {
	http.HandleFunc("/api/profile", middleware.ValidateSessionMiddleware(ProfileHandler))
}
