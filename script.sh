#!/bin/bash

# Ensure the script stops Docker containers on exit
trap cleanup EXIT

# Function to stop Docker containers
cleanup() {
    echo "Stopping Docker containers..."
    cd internal/database || { echo "Directory internal/database not found"; exit 1; }
    docker-compose down
    echo "Docker containers stopped."
}

# Navigate to the directory containing the Docker Compose file
cd internal/database || { echo "Directory internal/database not found"; exit 1; }

# Start Docker containers
echo "Starting Docker containers..."
docker-compose up -d
if [ $? -ne 0 ]; then
    echo "Failed to start Docker containers"
    exit 1
fi
echo "Docker containers started."

# Navigate back to the project root directory to run the Go application
cd ../../ || { echo "Failed to navigate to project root directory"; exit 1; }

# Run Go application
echo "Running Go application..."
go run cmd/socialNetwork/main.go
if [ $? -ne 0 ]; then
    echo "Go application exited with an error"
    exit 1
fi
echo "Go application exited successfully."