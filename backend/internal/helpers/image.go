package helpers

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
)

var (
	serialFilePath = "internal/database/images/serial.txt"
	serialLock     sync.Mutex
)

// SaveBase64Image decodes the Base64 image and saves it with a unique code as the name.
func SaveBase64Image(base64Image string) (string, error) {

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

	// Get the image format
	extension := getImageFormat([]byte(imageData))
	if extension == "unknown format" {
		return "", fmt.Errorf("unknown image format")
	}

	// Define the file path with correct extension
	fileName := uniqueCode + "." + extension
	filePath := filepath.Join("internal/database/images", fileName)

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

// getImageFormat returns the format of the image data
func getImageFormat(imageData []byte) string {
	switch {
	case bytes.HasPrefix(imageData, []byte{0xff, 0xd8, 0xff}):
		return "jpg" // or jpeg
	case bytes.HasPrefix(imageData, []byte{0x89, 'P', 'N', 'G', 0x0d, 0x0a, 0x1a, 0x0a}):
		return "png"
	case bytes.HasPrefix(imageData, []byte{'G', 'I', 'F', '8', '7', 'a'}) || bytes.HasPrefix(imageData, []byte{'G', 'I', 'F', '8', '9', 'a'}):
		return "gif"
	case bytes.HasPrefix(imageData, []byte{'B', 'M'}):
		return "bmp"
	case bytes.HasPrefix(imageData, []byte{'I', 'I', '*', 0x00}) || bytes.HasPrefix(imageData, []byte{'M', 'M', 0x00, '*'}):
		return "tiff"
	case bytes.HasPrefix(imageData, []byte{'<', '?', 'x', 'm', 'l'}) || bytes.HasPrefix(imageData, []byte{'<', 's', 'v', 'g'}):
		return "svg"
	case bytes.HasPrefix(imageData, []byte{'R', 'I', 'F', 'F'}) && bytes.HasPrefix(imageData[8:], []byte{'W', 'E', 'B', 'P'}):
		return "webp"
	default:
		return "unknown format"
	}
}
