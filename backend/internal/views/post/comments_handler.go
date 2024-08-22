package post

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	database "social-network/internal/database/querys"
	"social-network/internal/helpers"
	"social-network/internal/models"
	"social-network/internal/views/middleware"
)

/*
CreateCommentHandler creates a new comment on a post.

This function creates a new comment associated with a particular post.
It requires a valid user session to create a comment.

Example:

	// To create a new comment on a post
	POST /api/post/{id}/comments

	{
	    "body": string,
	    "image_id": string // optional
	}
*/
func CreateCommentHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	var comment models.Create_Comment
	parentID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Println("Failed to convert post ID to int:", err)
		helpers.HTTPError(w, "Invalid post ID:", http.StatusBadRequest)
		return
	}
	comment.ParentID = parentID
	err = json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		log.Println("Failed to decode comment:", err)
		helpers.HTTPError(w, "Failed to decode comment:", http.StatusBadRequest)
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
		log.Println("Failed to create comment:", err)
		helpers.HTTPError(w, "Invalid post ID:", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "create comment successful")
}

/*
GetPostCommentsHandler retrieves comments for a specific post.

This function retrieves comments associated with a particular post ID.
It requires a valid user session to access the comments.

Example:

	// To retrieve comments for a post with ID 123
	GET api/posts/123/comments

Response:

		[
	    {
	        "comment_id": int,
	        "image": string // optional
					"body": string,
					"creation_date": "2024-07-08T16:10:12.595563Z", // unix timestamp
	        "commenter": {
	            "first_name": string,
	            "last_name": string,
	            "image": string, //optional
	            "user_name":string
	        }
	    }

]
*/
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
	comments, err := database.Get_PostComments_from_db(userID, postIDInt)
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
