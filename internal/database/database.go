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
		UserName  string
		Email     string
		Password  string
		FirstName string
		LastName  string
		Gender    string
	}{
		{"Alice", "alice@example.com", "password123", "Alice", "Smith", "Female"},
		{"Bob", "bob@example.com", "password123", "Bob", "Johnson", "Male"},
		{"Charlie", "charlie@example.com", "password123", "Charlie", "Brown", "Male"},
	}

	// Prepare the insert statement
	stmt := `INSERT INTO public.user (user_name, email, password, first_name, last_name, gender) 
					 VALUES ($1, $2, $3, $4, $5, $6)
					 ON CONFLICT (user_name, email) DO NOTHING`

	// Insert the dummy data
	for _, user := range dummyUsers {
		_, err := DB.Exec(context.Background(), stmt, user.UserName, user.Email, user.Password, user.FirstName, user.LastName, user.Gender)
		if err != nil {
			log.Fatalf("Error inserting dummy data: %v\n", err)
		}
	}

	log.Println("Dummy data inserted successfully")
}
