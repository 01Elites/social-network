package helpers

import (
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
)

var (
	serialFilePath = "/../database/images/serial.txt"
	serialLock     sync.Mutex
)

// saveBase64Image decodes the Base64 image and saves it with a unique code as the name.
func saveBase64Image(base64Image string) (string, error) {
	// Generate a unique serial number for the image
	uniqueCode, err := getUniqueNumber()
	if err != nil {
		return "", fmt.Errorf("failed to generate unique code: %w", err)
	}

	// Decode the Base64 string
	imageData, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64 image: %w", err)
	}

	// Define the file path
	fileName := uniqueCode + ".webp" // or any other appropriate extension
	filePath := filepath.Join("/../database/images", fileName)

	// Ensure the directory exists
	err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("failed to create directories: %w", err)
	}

	// Write the image to the file
	err = os.WriteFile(filePath, imageData, 0644)
	if err != nil {
		return "", fmt.Errorf("failed to write image to file: %w", err)
	}

	// Return the file name
	return fileName, nil
}

// getUniqueNumber generates a unique serial number for the image file name
func getUniqueNumber() (string, error) {
	serialLock.Lock()
	defer serialLock.Unlock()

	// Ensure the directory for the serial file exists
	err := os.MkdirAll(filepath.Dir(serialFilePath), os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("failed to create directories for serial file: %w", err)
	}

	// Open the serial file
	file, err := os.OpenFile(serialFilePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return "", fmt.Errorf("failed to open serial file: %w", err)
	}
	defer file.Close()

	// Read the current serial number
	var serial int64
	_, err = fmt.Fscanf(file, "%d", &serial)
	if err != nil && err != io.EOF {
		return "", fmt.Errorf("failed to read serial number: %w", err)
	}

	// Increment the serial number
	serial++

	// Seek to the start of the file and truncate it
	_, err = file.Seek(0, io.SeekStart)
	if err != nil {
		return "", fmt.Errorf("failed to seek to start of serial file: %w", err)
	}
	err = file.Truncate(0)
	if err != nil {
		return "", fmt.Errorf("failed to truncate serial file: %w", err)
	}

	// Write the new serial number back to the file
	_, err = fmt.Fprintf(file, "%d", serial)
	if err != nil {
		return "", fmt.Errorf("failed to write serial file: %w", err)
	}

	// Return the serial number as a zero-padded string
	return fmt.Sprintf("%06d", serial), nil
}

// getImage returns the image in 64bit format.
func getImage(fileName string) (string, error) {
	// Read the image file
	imageData, err := os.ReadFile(filepath.Join("/../database/images", fileName))
	if err != nil {
		return "", fmt.Errorf("failed to read image file: %w", err)
	}

	// Encode the image data as a Base64 string
	base64Image := base64.StdEncoding.EncodeToString(imageData)

	return base64Image, nil
}
