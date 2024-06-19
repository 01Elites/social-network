package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

var (
	DB *pgx.Conn
)

func Init() {
	// Load environment variables
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Validate environment variables
	if dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		log.Fatalf("One or more environment variables are not set")
	}

	// Create the PostgreSQL connection string
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	// Connect to the PostgreSQL database
	conn, err := pgx.Connect(context.Background(), psqlInfo)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	// Ping the database to ensure the connection is established
	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatalf("Unable to ping the database: %v\n", err)
	}

	// Assign the initialized database connection to the global variable
	DB = conn
}

func InsertDummyData() {
	// Define the dummy user data
	dummyUsers := []struct {
		UserID		string
		UserName  string
		Email     string
		Password  string
		Provider string
		}{
			{"123e4567-e89b-12d3-a456-426614174000", "Alice", "alice@example.com", "password123", "password"},
			{"123e4567-e89b-12d3-a456-426614174001", "Bob", "bob@example.com", "password123", "google"},
	}

	// Prepare the insert statement with the provider field
	stmt := `INSERT INTO public.user (user_id, user_name, email, password, provider) 
					 VALUES ($1, $2, $3, $4, $5)
					 ON CONFLICT (user_name, email) DO NOTHING`

	// Insert the dummy data
	for _, user := range dummyUsers {
		_, err := DB.Exec(context.Background(), stmt, user.UserID, user.UserName, user.Email, user.Password, user.Provider)
		if err != nil {
			log.Print("Error inserting dummy data: %v\n", err)
		}
	}

	log.Println("Dummy data inserted successfully")
}
