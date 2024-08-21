package main

import (
	"log"
	"net/http"
	"social-network/internal/database/querys"
	"social-network/internal/helpers"
	"social-network/internal/views"
)

func main() {
	// Load environment variables from .env file
	helpers.LoadEnv(".env")

	// Apply database migrations
	err := querys.ApplyMigrations()
	if err != nil {
		log.Fatalf("Could not apply migrations: %v", err)
	}

	// Initiate the database connection
	querys.Init()

	// Setup the routes for the views
	views.SetupRoutes()

	// Print a message indicating that the server is live
	log.Println("\033[1;33mServer is Live at http://localhost:8081...\033[0m")
	log.Fatalln(http.ListenAndServe(":8081", nil))
}
