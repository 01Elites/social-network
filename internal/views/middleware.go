package views

import (
	"net/http"

	"social-network/internal/database"
)

func validateSessionMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := validateSession(w, r); err != nil {
			http.Error(w, "Invalid session token", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

func validateSession(w http.ResponseWriter, r *http.Request) error {
	// Extract the session token from the cookie
	cookie, err := r.Cookie("SN_SESSION")
	if err != nil {
		return err // No cookie means no session
	}
	// Validate the session token in the database
	if _, err = database.ValidateSessionToken(cookie.Value); err != nil {
		return err // Invalid or expired session token
	}

	setSessionCookie(w, cookie.Value)
	return nil
}
