package models

// Providers is a struct that holds all possible values for ProviderType.
var Provider = struct {
	Google string
	Github string
	Manual string
	Reboot string
}{
	Google: "google",
	Github: "github",
	Manual: "manual",
	Reboot: "reboot",
}

var ProfilePrivacy = struct {
	Private string
	Public  string
}{
	Private: "private",
	Public:  "public",
}

var Gender = struct {
	Male   string
	Female string
}{
	Male:   "male",
	Female: "female",
}

// User struct defines the structure for storing user information
type User struct {
	UserID    string          `db:"user_id"`
	Email     string          `db:"email"`
	Password  string          `db:"password"`
	Provider  string          `db:"provider"`
	UserName  string
	Following map[string]bool `json:"following"`
	Groups    map[int]bool    `json:"groups"`
}

// UserProfile struct defines the structure for storing user profile information
type UserProfile struct {
	UserID      string `db:"user_id"`
	NickName    string `db:"nick_name"`
	FirstName   string `db:"first_name"`
	LastName    string `db:"last_name"`
	Gender      string `db:"gender"`
	DateOfBirth string `db:"date_of_birth"`
	Image       string `db:"image"`
	Type        string `db:"type"`
}
