package views 

import (
	// "encoding/json"
	"fmt"
	// "log"
	"net/http"
	// "strconv"

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
	post := events.Create_Post{"Test1", "To check feed", "public", 0}
	user := events.User{"123e4567-e89b-12d3-a456-426614174000", "Alice", "alice@example.com", "password123", "password"}
	err := database.CreatePostInDB(user.ID, post)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}
}

func GetPostsHandler(w http.ResponseWriter, r *http.Request){
	posts, _ := database.GetPostsFeed("123e4567-e89b-12d3-a456-426614174000")
	fmt.Println(posts)
}