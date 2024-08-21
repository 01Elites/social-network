package pic

import (
	"log"
	"net/http"
	"path/filepath"

	database "social-network/internal/database/querys"
	"social-network/internal/helpers"
	"social-network/internal/views/middleware"
)

// servesImage handles gets the path of the image from url and serves it in the front end  as []byte
func servesImage(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		log.Printf("Error extracting token. User ID not found\n")
		helpers.HTTPError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	// get the last part of the url
	fileName := r.URL.Path[len("/api/image/"):]
	isAuthorized, err := database.CanUserSeeImage(userID, fileName)
	if !isAuthorized {
		helpers.HTTPError(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	if err != nil {
		log.Printf("Error checking if user can see image: %v\n", err)
		helpers.HTTPError(w, "Something Went Wrong!!", http.StatusBadRequest)
		return
	}
	path := filepath.Join("internal/database/images", fileName)
	http.ServeFile(w, r, path)
}
