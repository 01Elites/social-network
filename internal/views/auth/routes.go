package auth

import (
	"net/http"
	"social-network/internal/views/middleware"
)

func SetupAuthRoutes() {
	http.HandleFunc("POST /api/auth/signup", middleware.AllowCORS(SignUp))
	http.HandleFunc("POST /api/auth/signin", middleware.AllowCORS(SignIn))
	http.HandleFunc("DELETE /api/auth/logout", middleware.AllowCORS(LogOut))
}
