package post

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	database "social-network/internal/database/querys"
	"social-network/internal/helpers"
	"social-network/internal/models"
	"social-network/internal/views/middleware"
)

/*
CreatePostHandler creates a new post.

This function creates a new post using the data provided in the request body.
It requires a valid user session to create a post.

Example:

	    // To create a new post
	    POST /api/post
	    Body: {
	    "title": string,
	    "body": string,
	    "image": string, // optional
	    "privacy": "private"| "public" | "almost_private",
	    "usernames": []string // only if privacy == "almost_private"
		}
*/
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
		helpers.HTTPError(w, "content cannot be empty", http.StatusBadRequest)
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

/*
GetPostsHandler A function that retrieve posts, (used for mainpage posts listing)
This function fetches posts based on the user session and on whether the user follows the poster.

Example:

	// To retrieve posts
	GET api/posts

Response:

[

	{
		"post_id": int,
		"poster": {
				"user_name": string,
				"first_name": string,
				"last_name": string
		},
		"title": string,
		"content": string,
		"creation_date": "2024-07-07T14:28:45.127591Z", // unix time
		"post_privacy": "private"| "public" | "almost_private",
		"likes_count": int,
		"comments_count": int,
		"likers_usernames": null | []string // optional
	}

]
*/
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

/*
GetPostByIDHandler retrieves a single post by its ID.
This function retrieves a post based on the provided post ID.
It requires a valid user session to access the post.

Example:

	// To retrieve a post with ID 123
	GET api/post/123

Response:

	{
	    "post_id": int,
	    "poster": {
	        "user_name": string,
	        "first_name": string,
	        "last_name": string
	    },
	    "title": string
	    "content": string
	    "creation_date": "2024-07-07T14:28:45.127591Z", // unix time
	    "likes_count": int,
	    "comments_count": int,
			"likers_usernames": null | []string // optional
	}
*/
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

/*
DeletePostHandler deletes a single post by its ID.
This function deletes a post based on the provided post ID.
It requires a valid user session to access the post.

Example:

	// To delete a post with ID 123
	DELETE api/posts/123

Response:

	   Success -> 200 OK
	}
*/
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
		if err.Error() == "user unauthorized" {
			helpers.HTTPError(w, "user is unauthorized to delete the post", http.StatusUnauthorized)
			return
		}
		log.Printf("Failed to delete post: %v\n", err)
		helpers.HTTPError(w, "invalid post id", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Success")
}
