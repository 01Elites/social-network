package post

import (
	"encoding/json"
	"net/http"

	database "social-network/internal/database/querys"
	"social-network/internal/helpers"
	"social-network/internal/models"
	"social-network/internal/views/middleware"
)

/*
CreateLikeHandler adds a like to a post.

This function creates a new like associated with a particular post.
It requires a valid user session to create a like.

Example:

	    // To create a new like on a post
	    POST /api/post/like
Body: {
    "id": 0,
     }
*/
func CreateLikeHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		helpers.HTTPError(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	var post models.ID
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		helpers.HTTPError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// chaek if the post_id exists
	count, err := database.GetPostCountByID(post.ID)
	if err != nil || count == 0 {
		helpers.HTTPError(w, ("Post not found" + err.Error()), http.StatusBadRequest)
		return
	}

	// Update the like in the database
	err = database.UpDateLikeInDB(userID, post.ID)
	if err != nil {
		helpers.HTTPError(w, ("Failed to create like:" + err.Error()), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
