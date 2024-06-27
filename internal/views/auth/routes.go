package auth

import "net/http"

func SetupAuthRoutes() {
	http.HandleFunc("POST /api/auth/signup", SignUp)
	http.HandleFunc("POST /api/auth/signin", SignIn)
	http.HandleFunc("DELETE /api/auth/logout", LogOut)
}
