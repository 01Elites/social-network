package views

import (
	"net/http"
	"social-network/internal/views/auth"
	"social-network/internal/views/group"
	"social-network/internal/views/notifications"
	"social-network/internal/views/middleware"
	"social-network/internal/views/post"
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

	// http.HandleFunc("/", RootHandler)
	// http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("www/static/"))))

	// /********************* Authentication ************************/
	auth.SetupAuthRoutes()

	/********************* User ************************/
	profile.SetupProfileRoutes()
	// http.HandleFunc("GET /api/whoami", WhoAmI) // Handle whoami

	// /********************* Posts ************************/
	post.SetupPostRoutes()

	http.HandleFunc("OPTIONS /api/", middleware.AllowCORS(func(w http.ResponseWriter, r *http.Request) {}))

	// /********************* Group ************************/
	group.SetupGroupRoutes()

	notifications.SetupNotificationRoutes()
	// /********************* Categories ************************/
	// http.HandleFunc("GET /api/stats", GetStatsHandler)
	// http.HandleFunc("GET /api/categories", GetCategoriesHandler)

	// /********************* Chat ************************/
	// http.HandleFunc("/ws", HandleWebSocket)
}
