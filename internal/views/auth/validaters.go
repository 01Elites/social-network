package auth

import (
	"errors"
	"strings"
)

// checkPassWS checks if the given password string contains any whitespace characters.
// It returns true if there are whitespace characters, and false otherwise.
func checkWS(pass string) bool {
	return strings.ContainsAny(pass, " \t\n\r\v\f")
}

func ValidatePassword(pass string) error {
	// Validate password
	if checkWS(pass) {
		return errors.New("password cannot contain whitespaces")
	}

	if len(pass) < 6 || len(pass) > 20 {
		return errors.New("password should be between 6 and 20 characters long")
	}
	return nil
}
