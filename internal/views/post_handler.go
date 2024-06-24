package views

import (
	"encoding/json"
	"fmt"
	"log"

	// "log"
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
	err = database.CreatePostInDB(userID, post)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}
}

func GetPostsHandler(w http.ResponseWriter, r *http.Request){
	userID, ok := r.Context().Value(userIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	user, err := database.GetUserByID(userID);if err != nil {
		log.Print(err)
	}
	user.Following, err = database.GetUsersFollowingByID(userID);if err != nil {
		log.Print(err)
	}
	user.Groups, err = database.GetUserGroups(userID);if err != nil {
		log.Print(err)
	}
	posts, _ := database.GetPostsFeed(*user)
	fmt.Println(posts)
}

func GetPostByIDHandler(w http.ResponseWriter, r *http.Request) {
	// user := ValidateSession(w, r)

	postID := r.PathValue("id")
	postIDInt, _ := strconv.Atoi(postID)
	if postIDInt == 0 {
		http.Error(w, "Invalid post ID", http.StatusBadRequest)
		return
	}
	post, err := database.GetPostByID(postIDInt)
	if err != nil {
		http.Error(w, "Failed to get post", http.StatusBadRequest)
		return
	}
	fmt.Println(post)
	// if err := database.InsertPostView(postIDInt,user.ID); err != nil {
	// 	log.Printf("Failed to insert post view: %v", err)
	// }
	// json.NewEncoder(w).Encode(post)
}