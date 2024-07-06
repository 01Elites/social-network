package post

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	database "social-network/internal/database/querys"
	"social-network/internal/helpers"
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
	parentID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		helpers.HTTPError(w, ("Failed to decode post id" + err.Error()), http.StatusBadRequest)
		return
	}
	comment.ParentID = parentID
	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		helpers.HTTPError(w, ("Failed to decode comment:" + err.Error()), http.StatusBadRequest)
		return
	}
	if comment.Image != "" && comment.Image != "null" {
		comment.Image, err = helpers.SaveBase64Image(comment.Image)
		if err != nil {
			log.Println("Error with Image:", err)
			return
		}
	}
	err = database.Create_Comment_in_db(userID, comment)
	if err != nil {
		helpers.HTTPError(w, ("Invalid post ID:" + err.Error()), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "create comment successful")
}

func GetPostCommentsHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	postIDInt, _ := strconv.Atoi(r.PathValue("id"))
	if postIDInt == 0 {
		helpers.HTTPError(w, "Invalid post ID", http.StatusBadRequest)
		return
	}
	pageStr := r.URL.Query().Get("page")
	page, _ := strconv.Atoi(pageStr)
	comments, err := database.Get_PostComments_from_db(userID, postIDInt, page)
	if err != nil {
		helpers.HTTPError(w, ("Invalid post ID:" + err.Error()), http.StatusNotFound)
		return
	}
	commentsCapsul := struct {
		CommentsFeed []models.Comment `json:"comments"`
	}{
		CommentsFeed: comments,
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(commentsCapsul)
}
