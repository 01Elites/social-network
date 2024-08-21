#!/bin/bash

# Make this script executable
chmod +x $0

# Ensure the script stops Docker containers on exit
trap cleanup EXIT

# Function to stop Docker containers
cleanup() {
	echo "Stopping Docker containers..."
	docker-compose down
	echo "Docker containers stopped."
}

# Start Docker containers
echo "Starting Docker containers..."
# docker pull postgres
# docker pull caddy
docker-compose up
