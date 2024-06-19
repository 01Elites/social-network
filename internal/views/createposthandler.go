package views 

import (
	"encoding/json"
	"fmt"
	// "log"
	"net/http"
	// "strconv"

	"social-network/internal/database"
	"social-network/internal/events"
)

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var post events.Create_Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "Failed to decode post", http.StatusBadRequest)
		return
	}
	err = database.CreatePostInDB(user.ID, post)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to create post", http.StatusInternalServerError)
		return
	}
}