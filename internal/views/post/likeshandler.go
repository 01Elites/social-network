package post

import (
	"encoding/json"
	"net/http"
	"strconv"

	database "social-network/internal/database/querys"
	"social-network/internal/helpers"
	"social-network/internal/views/middleware"
)

// CreateLikeHandler handles the creation of a like for posts & comments.
func CreateLikeHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		helpers.HTTPError(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	post_id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helpers.HTTPError(w, ("Invalid postID:" + err.Error()), http.StatusBadRequest)
		return
	}

	// chaek if the post_id exists
	count, err := database.GetPostCountByID(post_id)
	if err != nil || count == 0 {
		helpers.HTTPError(w, ("Post not found" + err.Error()), http.StatusBadRequest)
		return
	}

	// Update the like in the database
	err = database.UpDateLikeInDB(userID, post_id)
	if err != nil {
		helpers.HTTPError(w, ("Failed to create like:" + err.Error()), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post_id)
}
