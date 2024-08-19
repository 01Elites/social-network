package views

import (
	"net/http"

	"social-network/internal/views/auth"
	"social-network/internal/views/follow"
	"social-network/internal/views/friends"
	"social-network/internal/views/group"
	"social-network/internal/views/middleware"
	"social-network/internal/views/pic"
	"social-network/internal/views/post"
	"social-network/internal/views/profile"
	"social-network/internal/views/websocket"
	"social-network/internal/views/events"
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

	http.HandleFunc("/api/", middleware.AllowCORS(func(w http.ResponseWriter, r *http.Request) {}))

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

	/************************* Image *****************************/
	pic.SetupImageRoutes()

	/*********************** Websocket ***************************/
	websocket.SetupWebSocketRoutes()

	/************************ Friends ****************************/
	friends.SetupFriendsRoutes()

	/************************ Events ****************************/
	events.SetupEventsRoutes()
}
