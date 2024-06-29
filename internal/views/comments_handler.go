package views

import (
	"encoding/json"
	"io"
	"net/http"
	"social-network/internal/database"
	"social-network/internal/models"
	"strconv"
)

func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(userIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	var comment models.Create_Comment
	parentID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Failed to decode post id", http.StatusBadRequest)
		return
	}
	comment.ParentID = parentID
	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonError := models.Error{
				Reason: "failed to decode comment",
		}
		json.NewEncoder(w).Encode(jsonError)
		return
	}
	err = database.Create_Comment_in_db(userID, comment)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		jsonError := models.Error{
				Reason: "invalid post id",
		}
		json.NewEncoder(w).Encode(jsonError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "create comment successful")
}

func GetPostCommentsHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(userIDKey).(string)
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
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "get comments successful")
}
