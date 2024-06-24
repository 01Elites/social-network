package views 

import (
	// "encoding/json"
	"fmt"
	// "log"
	"net/http"
	"strconv"

	"social-network/internal/database"
	"social-network/internal/events"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	// err := json.NewDecoder(r.Body).Decode(&post)
	// if err != nil {
	// 	http.Error(w, "Failed to decode post", http.StatusBadRequest)
	// 	return
	// }
	fmt.Println("Stestse")
	post := events.Create_Post{"Test1", "To check feed", "public", 1}
	user := events.User{"123e4567-e89b-12d3-a456-426614174000", "Alice", "alice@example.com", "password123", "password", nil, nil}
	err := database.CreatePostInDB(user.ID, post)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}
}

func GetPostsHandler(w http.ResponseWriter, r *http.Request){
	var dummyUser events.User
	posts, _ := database.GetPostsFeed(dummyUser)
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