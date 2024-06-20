package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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
		UserID   string
		UserName string
		Email    string
		Password string
		Provider string
	}{
		{"123e4567-e89b-12d3-a456-426614174000", "Alice", "alice@example.com", "password123", "password"},
		{"123e4567-e89b-12d3-a456-426614174001", "Bob", "bob@example.com", "password123", "google"},
		{"123e4567-e89b-12d3-a456-426614174002", "Charlie", "charlie@example.com", "password123", "github"},
	}

	// Prepare the insert statement with the provider field
	stmt := `INSERT INTO public.user (user_id, user_name, email, password, provider) 
					 VALUES ($1, $2, $3, $4, $5)
					 ON CONFLICT (user_name) DO NOTHING`

	// Insert the dummy data
	for _, user := range dummyUsers {
		_, err := DB.Exec(context.Background(), stmt, user.UserID, user.UserName, user.Email, user.Password, user.Provider)
		if err != nil {
			log.Fatalf("Error inserting dummy data: %v\n", err)
		}
	}

	log.Println("Dummy data inserted successfully")
}

func ApplyMigrations() error {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Create the PostgreSQL connection string
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, dbPassword, dbHost, dbPort, dbName)

	m, err := migrate.New(
		"file://internal/database/migrations",
		dbURL)
	if err != nil {
		fmt.Println("Error in migration")
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		fmt.Println("Error in migration")
		return err
	}
	log.Println("Migrations applied successfully")
	return nil
}
