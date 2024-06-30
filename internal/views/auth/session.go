package auth

import (
	"net/http"
	"time"
)

func clearSessionCookie(w http.ResponseWriter) {
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
