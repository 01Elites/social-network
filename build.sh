#!/bin/bash

# Make this script executable
chmod +x $0

# Navigate to the directory containing the Docker Compose file
cd internal/database

# Start Docker containers
sudo docker-compose up --build -d

# Navigate back to the project root directory to run the Go application
cd ../../

# Run Go application
go run cmd/socialNetwork/main.go 

# When the Go application exits, stop Docker containers
cd internal/database
sudo docker-compose down
