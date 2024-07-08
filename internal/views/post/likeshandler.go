package post

import (
	"io"
	"log"
	"net/http"
	"strconv"

	database "social-network/internal/database/querys"
	"social-network/internal/helpers"
	"social-network/internal/views/middleware"
)

/*
CreateLikeHandler adds a like to a post.

This function creates a new like associated with a particular post.
It requires a valid user session to create a like.

Example:

	// To create a new like on a post
	POST /api/post/{id}/like
*/
func CreateLikeHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		helpers.HTTPError(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	postID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Printf("Failed to decode post id: %v\n", err)
		helpers.HTTPError(w, "Failed to decode post id", http.StatusBadRequest)
		return
	}

	// chaek if the post_id exists
	exists, err := database.PostExists(postID)
	if err != nil {
		log.Printf("Failed to check if post exists: %v\n", err)
		helpers.HTTPError(w, "Failed to check if post exists:", http.StatusInternalServerError)
		return
	}
	if !exists {
		helpers.HTTPError(w, "Post does not exist", http.StatusNotFound)
		return
	}

	// Update the like in the database
	err = database.UpDateLikeInDB(userID, postID)
	if err != nil {
		log.Printf("Failed to create like: %v\n", err)
		helpers.HTTPError(w, "Failed to create like:", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Like created successfully")
}
