package session

import (
	"net/http"
	"time"

	"social-network/internal/database"
)

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

	SetSessionCookie(w, cookie.Value)
	return userID, nil
}

func ClearSessionCookie(w http.ResponseWriter) {
	// Clear the session cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "SN_SESSION",
		Value:    "",
		Expires:  time.Now().Add(-10 * time.Second), // to enhance the compatibility and ensure all browsers handle the cookie clearing
		Path:     "/",
		MaxAge:   -1,
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
	})
}

// updateSessionCookie updates the session cookie expiration time.
func SetSessionCookie(w http.ResponseWriter, sessionToken string) {
	expiration := time.Now().AddDate(1, 0, 0)
	updatedCookie := http.Cookie{
		Name:     "SN_SESSION",
		Value:    sessionToken,
		Expires:  expiration,
		Path:     "/",
		SameSite: http.SameSiteStrictMode,
		Secure:   true,
		HttpOnly: true,
	}
	http.SetCookie(w, &updatedCookie)
}
