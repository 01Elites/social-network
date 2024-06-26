package views

import (
	"context"
	"net/http"

	"social-network/internal/database"
)

type contextKey string
		
const userIDKey contextKey = "userID"

func validateSessionMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID, err := validateSession(w, r)
		if err != nil {
			http.Error(w, "Invalid session token", http.StatusUnauthorized)
			return
		}
		
		ctx := context.WithValue(r.Context(), userIDKey, userID)
		next(w, r.WithContext(ctx))
	}
}

func validateSession(w http.ResponseWriter, r *http.Request) (string, error) {
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

	setSessionCookie(w, cookie.Value)
	return userID, nil
}
