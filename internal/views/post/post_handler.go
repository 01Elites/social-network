package post

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"social-network/internal/database"
	"social-network/internal/helpers"
	"social-network/internal/models"
	"social-network/internal/views/middleware"
	"strconv"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	var post models.Create_Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil || post.Title == "" || post.Content == "" {
		w.WriteHeader(http.StatusBadRequest)
		jsonError := models.Error{
			Reason: "failed to decode post",
		}
		json.NewEncoder(w).Encode(jsonError)
		return
	}
	if post.Image != "" {
		post.Image, err = helpers.SaveBase64Image(post.Image)
		if err != nil {
			fmt.Println("Error with Image:\n", err)
		}
	}

	postID, err := database.CreatePostInDB(userID, post)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		jsonError := models.Error{
			Reason: err.Error(),
		}
		json.NewEncoder(w).Encode(jsonError)
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
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	user, err := database.GetUserByID(userID)
	if err != nil {
		log.Print(err)
	}
	user.Following, err = database.GetUsersFollowingByID(userID)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to get followers", http.StatusInternalServerError)
		return
	}
	posts, err := database.GetPostsFeed(*user)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to get posts", http.StatusInternalServerError)
		return
	}
	fmt.Println(posts)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func GetPostByIDHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	postID := r.PathValue("id")
	postIDInt, err := strconv.Atoi(postID)
	if postIDInt == 0 || err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonError := models.Error{
			Reason: "invalid post id",
		}
		json.NewEncoder(w).Encode(jsonError)
		return
	}
	post, err := database.GetPostByID(postIDInt, userID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		jsonError := models.Error{
			Reason: "invalid post id",
		}
		json.NewEncoder(w).Encode(jsonError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}

func DeletePostHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	postID := r.PathValue("id")
	postIDInt, err := strconv.Atoi(postID)
	if postIDInt == 0 || err != nil {
		w.WriteHeader(http.StatusNotFound)
		jsonError := models.Error{
			Reason: "invalid post id",
		}
		json.NewEncoder(w).Encode(jsonError)
		return
	}
	err = database.DeletePost(postIDInt, userID)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
}
