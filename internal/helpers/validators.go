package helpers

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"social-network/internal/models"
)

// validateAge checks if the age calculated from the given dob is less than 18.
func ValidateAge(dob time.Time) error {
	// Calculate the difference in years between the current time and the dob.
	currentTime := time.Now()
	age := currentTime.Year() - dob.Year()

	// Adjust the age if the current date is before the birthday in the current year.
	if currentTime.Month() < dob.Month() || (currentTime.Month() == dob.Month() && currentTime.Day() < dob.Day()) {
		age--
	}

	if age < 18 {
		return errors.New("you must be at least 18 years old to sign up")
	}

	return nil
}

// ValidateEmail checks if the email address is valid
func ValidateEmail(email *string) error {
	*email = strings.TrimSpace(*email)
	if *email == "" {
		return errors.New("email is required")
	}
	// Define the regular expression for a valid email address
	const emailRegexPattern = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	// Compile the regular expression
	emailRegex := regexp.MustCompile(emailRegexPattern)
	// Match the email with the regular expression
	if !emailRegex.MatchString(*email) {
		return errors.New("invalid email address")
	}
	return nil
}

func ValidateUnemptyFields(data *models.UserProfile) error {
	// Trim spaces and check for empty required fields
	data.FirstName = strings.TrimSpace(data.FirstName)
	data.LastName = strings.TrimSpace(data.LastName)

	if data.FirstName == "" || data.LastName == "" {
		return errors.New("all fields must be filled out")
	}

	if data.DateOfBirth.IsZero() {
		return errors.New("date of birth is required")
	}

	if data.Gender == "" {
		return errors.New("gender is required")
	}

	return nil
}

// ValidateSignUpData validates the sign-up data and returns an error if any validation fails.
func ValidateUserProfileData(data *models.UserProfile) error {
	// Check the length of the username
	if len(data.NickName) > 20 {
		return errors.New("nickname should be up to 20 characters long")
	}

	// Check the length of the first name and last name
	if len(data.FirstName) > 20 || len(data.LastName) > 20 {
		return errors.New("first name and last name should be up to 20 characters long")
	}
	// Validate Date of Birth to check if user is at least 18 years old
	if err := ValidateAge(data.DateOfBirth); err != nil {
		return err
	}

	// Validate About field
	if len(data.About) > 500 {
		return errors.New("about section should be up to 500 characters long")
	}
	// Validate Gender
	switch data.Gender {
	case models.Gender.Female, models.Gender.Male:
		// valid
	default:
		return fmt.Errorf("invalid gender: %v", data.Gender)
	}

	// Validate privacy
	if data.ProfilePrivacy != "" {
		switch data.ProfilePrivacy {
		case models.ProfilePrivacy.Private, models.ProfilePrivacy.Public:
			// valid
		default:
			return fmt.Errorf("invalid profile privacy: %v", data.ProfilePrivacy)
		}
	}

	return nil
}
