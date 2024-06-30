package views

import (
	"net/http"
	"social-network/internal/views/auth"
	"social-network/internal/views/middleware"
	"social-network/internal/views/profile"
)

// SetupRoutes sets up the routes for the web application.
func SetupRoutes() {
	// test Handle the root route and return a simple message "postman"
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api" {
			http.NotFound(w, r)
			return
		}
		w.Write([]byte("Hello, World!"))
	})

	m := middleware.ValidateSessionMiddleware
	// http.HandleFunc("/", RootHandler)
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("www/static/"))))

	// /********************* Authentication ************************/
	auth.SetupAuthRoutes()

	/********************* User ************************/
	profile.SetupProfileRoutes()
	// http.HandleFunc("GET /api/whoami", WhoAmI) // Handle whoami

	// /********************* Posts ************************/
	http.HandleFunc("/create_group", m(CreateGroupHandler))
	http.HandleFunc("/posts", m(GetPostsHandler))
	http.HandleFunc("GET /api/post/{id}", m(GetPostByIDHandler))
	http.HandleFunc("GET /api/post/{id}/comments", m(GetPostCommentsHandler))
	http.HandleFunc("/create_post", m(CreatePostHandler))
	http.HandleFunc("/create_comment", m(CreateCommentHandler))
	// http.HandleFunc("POST /api/create_like/{post_id}", CreateLikeHandler)

	// /********************* Categories ************************/
	// http.HandleFunc("GET /api/stats", GetStatsHandler)
	// http.HandleFunc("GET /api/categories", GetCategoriesHandler)

	// /********************* Chat ************************/
	// http.HandleFunc("/ws", HandleWebSocket)
}
