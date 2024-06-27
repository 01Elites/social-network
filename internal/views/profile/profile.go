package profile

import (
	"fmt"
	"net/http"
	"social-network/internal/views/middleware"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	// Retrieve the userID from context using the same key defined globally
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Profile of user: %s", userID)
}
