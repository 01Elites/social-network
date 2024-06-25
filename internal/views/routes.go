package views

import "net/http"

// SetupRoutes sets up the routes for the web application.
func SetupRoutes() {
	// test Handle the root route and return a simple message "postman"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Write([]byte("Hello, World!"))
	})
	// http.HandleFunc("/", RootHandler)
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("www/static/"))))

	// /********************* Authentication ************************/
	http.HandleFunc("POST /api/auth/signup", SignUp)
	http.HandleFunc("POST /api/auth/signin", SignIn)
	http.HandleFunc("DELETE /api/auth/logout", LogOut)

	/********************* User ************************/
	http.HandleFunc("GET /api/profile", validateSessionMiddleware(ProfileHandler))
	// http.HandleFunc("GET /api/whoami", WhoAmI) // Handle whoami

	// /********************* Posts ************************/
	http.HandleFunc("GET /api/posts", validateSessionMiddleware(GetPostsHandler))
	http.HandleFunc("GET /api/post/{id}", validateSessionMiddleware(GetPostByIDHandler))
	http.HandleFunc("GET /api/post/{id}/comments", validateSessionMiddleware(GetPostCommentsHandler))
	http.HandleFunc("POST /api/create_post", validateSessionMiddleware(CreatePostHandler))
	http.HandleFunc("POST /api/create_comment", validateSessionMiddleware(CreateCommentHandler))

	// /********************* Group ************************/
	http.HandleFunc("POST /api/create_group", validateSessionMiddleware(CreateGroupHandler))
	// http.HandleFunc("/invite_user", m(InvitationHandler))
	// http.HandleFunc("/group_request", m(RequestHandler))
	// http.HandleFunc("/search_group", m(SearchGroupHandler))
	// http.HandleFunc("/create_event", m(CreateEventHandler))
	// http.HandleFunc("/event_response", m(EventResponseHandler))
	// http.HandleFunc("POST /api/create_like/{post_id}", CreateLikeHandler)

	// /********************* Categories ************************/
	// http.HandleFunc("GET /api/stats", GetStatsHandler)
	// http.HandleFunc("GET /api/categories", GetCategoriesHandler)

	// /********************* Chat ************************/
	// http.HandleFunc("/ws", HandleWebSocket)
}
