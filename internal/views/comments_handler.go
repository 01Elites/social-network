package views

import (
	"encoding/json"
	"net/http"
	"social-network/internal/database"
	"social-network/internal/models"
	"social-network/internal/views/middleware"
	"strconv"
)

func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	var comment models.Create_Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, "Failed to decode comment", http.StatusBadRequest)
		return
	}
	err = database.Create_Comment_in_db(userID, comment)
	if err != nil {
		http.Error(w, "Failed to create comment", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func GetPostCommentsHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	postIDInt, _ := strconv.Atoi(r.PathValue("id"))
	if postIDInt == 0 {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}
	pageStr := r.URL.Query().Get("page")
	page, _ := strconv.Atoi(pageStr)
	comments, err := database.Get_PostComments_from_db(userID, postIDInt, page)
	if err != nil {
		http.Error(w, "Failed to get comment", http.StatusBadRequest)
		return
	}
	commentsCapsul := struct {
		CommentsFeed []models.Comment `json:"comments"`
	}{
		CommentsFeed: comments,
	}
	json.NewEncoder(w).Encode(commentsCapsul)
}
