package views

// SetupRoutes sets up the routes for the web application.
func SetupRoutes() {
	// http.HandleFunc("/", RootHandler)
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("www/static/"))))

	// /********************* Authentication ************************/
	// http.HandleFunc("POST /signup", SignUp)    // Handle signup
	// http.HandleFunc("POST /login", LogIn)      // Handle login
	// http.HandleFunc("DELETE /logout", LogOut)  //  Handle logout
	// http.HandleFunc("GET /api/whoami", WhoAmI) // Handle whoami

	// /********************* Posts ************************/
	// http.HandleFunc("GET /api/posts", GetPostsHandler)
	// http.HandleFunc("GET /api/post/{id}", GetPostByIDHandler)
	// http.HandleFunc("GET /api/post/{id}/comments", GetPostCommentsHandler)
	// http.HandleFunc("POST /api/create_post", CreatePostHandler)
	// http.HandleFunc("POST /api/create_comment", CreateCommentHandler)
	// http.HandleFunc("POST /api/create_like/{post_id}", CreateLikeHandler)

	// /********************* Categories ************************/
	// http.HandleFunc("GET /api/stats", GetStatsHandler)
	// http.HandleFunc("GET /api/categories", GetCategoriesHandler)

	// /********************* Chat ************************/
	// http.HandleFunc("/ws", HandleWebSocket)
}
