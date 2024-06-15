package helpers

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// LoadEnv reads a .env file and sets the environment variables
func LoadEnv(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening .env file: %v\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Ignore comments and empty lines
		if strings.HasPrefix(line, "#") || line == "" {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			log.Fatalf("Invalid line in .env file: %s\n", line)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// Set the environment variable
		err = os.Setenv(key, value)
		if err != nil {
			log.Fatalf("Error setting environment variable: %v\n", err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading .env file: %v\n", err)
	}
}
