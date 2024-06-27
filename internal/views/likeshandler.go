package views

import (
	"fmt"
	"net/http"
	"social-network/internal/database"
	"strconv"
	"encoding/json"
)

// CreateLikeHandler handles the creation of a like for posts & comments.
func CreateLikeHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(userIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	post_id, err := strconv.Atoi(r.PathValue("post_id"))
	if err != nil {
		http.Error(w, "Invalid post id", http.StatusBadRequest)
		return
	}

	//chaek if the post_id exists
	count, err := database.GetPostCountByID(post_id)
	if err != nil || count == 0 {
		http.Error(w, "Post not found", http.StatusBadRequest)
		return
	}

	// Update the like in the database
	err = database.UpDateLikeInDB(userID, post_id)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to create like", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(post_id)
}