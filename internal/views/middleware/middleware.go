package middleware

import (
	"context"
	"net/http"

	"social-network/internal/database"
	"social-network/internal/views/auth"
)

type contextKey string

const UserIDKey contextKey = "userID"

func ValidateSessionMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		userID, err := ValidateSession(w, r)
		if err != nil {
			http.Error(w, "Invalid session token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, userID)
		next(w, r.WithContext(ctx))
	}
}

func ValidateSession(w http.ResponseWriter, r *http.Request) (string, error) {
	// Extract the session token from the cookie
	cookie, err := r.Cookie("SN_SESSION")
	if err != nil {
		return "", err // No session token, user is not logged in
	}
	// Validate the session token in the database
	userID, err := database.ValidateSessionToken(cookie.Value)
	if err != nil {
		return "", err
	}

	auth.SetSessionCookie(w, cookie.Value)
	return userID, nil
}
