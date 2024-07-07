package views

import (
	"net/http"
	"social-network/internal/views/auth"
	"social-network/internal/views/follow"
	"social-network/internal/views/group"
	"social-network/internal/views/middleware"
	"social-network/internal/views/post"
	"social-network/internal/views/profile"
	"social-network/internal/views/websocket"
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

	// test Handle the OPTIONS request for the root route and return the allowed methods
	http.HandleFunc("OPTIONS /api/", middleware.AllowCORS(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	}))

	/********************* Authentication ************************/
	auth.SetupAuthRoutes()

	/************************** User *****************************/
	profile.SetupProfileRoutes()

	/*************************** Posts ***************************/
	post.SetupPostRoutes()

	/*************************** Group ***************************/
	group.SetupGroupRoutes()

	/************************** Follow ***************************/
	follow.SetupFollowRoutes()

	/*********************** Websocket ***************************/
	websocket.SetupWebSocketRoutes()
}
