package views

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"social-network/internal/database"
	"social-network/internal/models"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(userIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	var post models.Create_Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Failed to decode post", http.StatusBadRequest)
		return
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
	userID, ok := r.Context().Value(userIDKey).(string)
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
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
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
	// user := ValidateSession(w, r)

	postID := r.PathValue("id")
	postIDInt, _ := strconv.Atoi(postID)
	if postIDInt == 0 {
		w.WriteHeader(http.StatusBadRequest)
		jsonError := models.Error{
				Reason: "invalid post id",
		}
		json.NewEncoder(w).Encode(jsonError)
		return
	}
	post, err := database.GetPostByID(postIDInt)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		jsonError := models.Error{
				Reason: "invalid post id",
		}
		json.NewEncoder(w).Encode(jsonError)
		return
	}
	json.NewEncoder(w).Encode(post)
	w.WriteHeader(http.StatusOK)
}
