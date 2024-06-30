package auth

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"social-network/internal/models"
)

// ValidateSignUpData validates the sign-up data and returns an error if any validation fails.
func ValidateSignUpData(data *SignUpRequst) error {
	// Trim spaces and check for empty required fields
	data.Email = strings.TrimSpace(data.Email)
	data.NickName = strings.TrimSpace(data.NickName)
	data.FirstName = strings.TrimSpace(data.FirstName)
	data.LastName = strings.TrimSpace(data.LastName)

	if data.Email == "" || data.NickName == "" || data.FirstName == "" || data.LastName == "" {
		return errors.New("all fields must be filled out")
	}

	// Check the length of the username
	if len(data.NickName) > 20 {
		return errors.New("nickname should be up to 20 characters long")
	}

	// Check the length of the first name and last name
	if len(data.FirstName) > 20 || len(data.LastName) > 20 {
		return errors.New("first name and last name should be up to 20 characters long")
	}

	// Check the length of the email
	if len(data.Email) > 30 {
		return errors.New("email should be up to 30 characters long")
	}

	// Validate password
	if checkPassWS(data.Password) {
		return errors.New("password cannot contain whitespaces")
	}

	if len(data.Password) < 6 || len(data.Password) > 20 {
		return errors.New("password should be between 6 and 20 characters long")
	}

	// Validate Date of Birth to check if user is at least 18 years old
	if err := validateAge(data.DateOfBirth); err != nil {
		return err
	}

	// Validate privacy
	switch data.ProfilePrivacy {
	case models.ProfilePrivacy.Private, models.ProfilePrivacy.Public:
		// valid
	default:
		return fmt.Errorf("invalid profile privacy: %v", data.ProfilePrivacy)
	}

	return nil
}

// checkPassWS checks if the given password string contains any whitespace characters.
// It returns true if there are whitespace characters, and false otherwise.
func checkPassWS(pass string) bool {
	return strings.ContainsAny(pass, " \t\n\r\v\f")
}

// validateAge checks if the user is at least 18 years old based on their date of birth.
func validateAge(dob string) error {
	// Parse the date of birth string into a time.Time object
	dateOfBirth, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return fmt.Errorf("invalid date format: %v", err)
	}

	// Get the current date
	now := time.Now()

	// Calculate the user's age
	age := now.Year() - dateOfBirth.Year()
	if now.YearDay() < dateOfBirth.YearDay() {
		age--
	}

	// Check if the user is at least 18 years old
	if age < 18 {
		return errors.New("you must be at least 18 years old to sign up")
	}

	return nil
}
