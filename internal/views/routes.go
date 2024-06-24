package views

import "net/http"

// SetupRoutes sets up the routes for the web application.
func SetupRoutes() {
	// test Handle the root route and return a simple message "postman"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	// http.HandleFunc("/", RootHandler)
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("www/static/"))))

	// /********************* Authentication ************************/
	// http.HandleFunc("POST /signup", SignUp)    // Handle signup
	// http.HandleFunc("POST /login", LogIn)      // Handle login
	// http.HandleFunc("DELETE /logout", LogOut)  //  Handle logout
	// http.HandleFunc("GET /api/whoami", WhoAmI) // Handle whoami

	// /********************* Posts ************************/
	http.HandleFunc("/create_group", CreateGroupHandler)
	http.HandleFunc("/posts", GetPostsHandler)
	http.HandleFunc("GET /api/post/{id}", GetPostByIDHandler)
	// http.HandleFunc("GET /api/post/{id}/comments", GetPostCommentsHandler)
	http.HandleFunc("/create_post", CreatePostHandler)
	http.HandleFunc("/create_comment", CreateCommentHandler)
	// http.HandleFunc("POST /api/create_like/{post_id}", CreateLikeHandler)

	// /********************* Categories ************************/
	// http.HandleFunc("GET /api/stats", GetStatsHandler)
	// http.HandleFunc("GET /api/categories", GetCategoriesHandler)

	// /********************* Chat ************************/
	// http.HandleFunc("/ws", HandleWebSocket)
}
