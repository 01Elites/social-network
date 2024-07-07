package post

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	database "social-network/internal/database/querys"
	"social-network/internal/helpers"
	"social-network/internal/models"
	"social-network/internal/views/middleware"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		log.Printf("User ID not found in post handler\n")
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	var post models.Create_Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		log.Printf("Failed to decode post: %v\n", err)
		helpers.HTTPError(w, "invalid post", http.StatusBadRequest)
		return
	}

	// validate the post content
	post.Content = strings.TrimSpace(post.Content)
	if post.Content == "" {
		log.Printf("Post content is empty\n")
		helpers.HTTPError(w, "invalid post", http.StatusBadRequest)
		return
	}
	// save the image if it exists
	if post.Image != "" && post.Image != "null" {
		post.Image, err = helpers.SaveBase64Image(post.Image)
		if err != nil {
			fmt.Println("Error with Image:\n", err)
		}
	}

	if post.Privacy == "almost_private" && (post.UserNames == nil || len(post.UserNames) == 0) {
		helpers.HTTPError(w, "you must choose at least one user.", http.StatusBadRequest)
		return
	}
	postID, err := database.CreatePostInDB(userID, post)
	if err != nil {
		if strings.Contains(err.Error(), "(SQLSTATE 23503)") && strings.Contains(err.Error(), "post_group_id_fkey") {
			log.Printf("Failed to create post: %v\n", err)
			helpers.HTTPError(w, "invalid group id", http.StatusBadRequest)
			return
		}
		log.Printf("Failed to create post: %v\n", err)
		helpers.HTTPError(w, "invalid post data", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	postIDjson := struct {
		ID int `json:"id"`
	}{
		ID: postID,
	}
	json.NewEncoder(w).Encode(postIDjson)
}

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		log.Printf("User ID not found in post handler\n")
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	user, err := database.GetUserByID(userID)
	if err != nil {
		log.Printf("Failed to get user details: %v\n", err)
		helpers.HTTPError(w, "Failed to get userID", http.StatusInternalServerError)
		return
	}
	user.Following, err = database.GetUsersFollowingByID(userID)
	if err != nil {
		log.Printf("Failed to get following: %v\n", err)
		helpers.HTTPError(w, "Failed to get followers", http.StatusInternalServerError)
		return
	}
	posts, err := database.GetPostsFeed(*user)
	if err != nil {
		log.Printf("Failed to get posts: %v\n", err)
		helpers.HTTPError(w, "Failed to get posts", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func GetPostByIDHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		log.Printf("User ID not found in post handler\n")
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	postID := r.PathValue("id")
	postIDInt, err := strconv.Atoi(postID)
	if postIDInt == 0 || err != nil {
		helpers.HTTPError(w, "invalid post id", http.StatusNotFound)
		return
	}
	post, err := database.GetPostByID(postIDInt, userID)
	if err != nil {
		log.Printf("Failed to get post: %v\n", err)
		helpers.HTTPError(w, "invalid post id", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		log.Printf("User ID not found in post handler\n")
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	postID := r.PathValue("id")
	postIDInt, err := strconv.Atoi(postID)
	if postIDInt == 0 || err != nil {
		log.Printf("Failed to get post: %v\n", err)
		helpers.HTTPError(w, "invalid post id", http.StatusNotFound)
		return
	}
	err = database.DeletePost(postIDInt, userID)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
}
