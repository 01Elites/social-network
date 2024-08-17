package auth

import (
	"net/http"
	"social-network/internal/views/middleware"
)

func SetupAuthRoutes() {
	http.HandleFunc("POST /api/auth/signup", middleware.AllowCORS(SignUp))
	http.HandleFunc("POST /api/auth/signin", middleware.AllowCORS(SignIn))
	http.HandleFunc("DELETE /api/auth/logout", middleware.AllowCORS(LogOut))
	http.HandleFunc("/api/auth/gitea/login", middleware.AllowCORS(GiteaLogin))
	http.HandleFunc("/api/auth/gitea/callback", middleware.AllowCORS(GiteaCallback))
	http.HandleFunc("/api/auth/github/login", middleware.AllowCORS(HandleGithubLogin))
	http.HandleFunc("/api/auth/github/callback", middleware.AllowCORS(HandleGithubCallback))
	http.HandleFunc("/api/auth/google/login", middleware.AllowCORS(HandleGoogleLogin))
	http.HandleFunc("/api/auth/google/callback", middleware.AllowCORS(HandleGoogleCallback))
}
