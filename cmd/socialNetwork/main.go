package main

import (
	"log"
	"net/http"
	"social-network/backend/database"
	"social-network/backend/helpers"
	"social-network/backend/views"
)

func main() {
	// Load environment variables from .env file
	helpers.LoadEnv("backend/database/.env")

	// Initiate the database connection
	database.Init()
	database.InsertDummyData() // Insert dummy data for testing

	// Setup the routes for the views
	views.SetupRoutes()

	// Print a message indicating that the server is live
	log.Println("\033[1;33mServer is Live at http://localhost:8080...\033[0m")
	log.Fatalln(http.ListenAndServe(":8080", nil))
}
