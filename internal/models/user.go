package models

import "time"

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
	UserName  string          `db:"user_name"`
	Email     string          `db:"email"`
	Password  string          `db:"password,omitempty"`
	Provider  string          `db:"provider"`
	Following map[string]bool `json:"following"`
	// Groups    map[int]bool    `json:"groups"`
}

// UserProfile struct defines the structure for storing user profile information
type UserProfile struct {
	UserID         string    `db:"user_id"`
	NickName       string    `db:"nick_name"`
	FirstName      string    `db:"first_name"`
	LastName       string    `db:"last_name"`
	Gender         string    `db:"gender"`
	DateOfBirth    time.Time `db:"date_of_birth"`
	ProfilePrivacy string    `db:"privacy"`
	Image          string    `db:"image"`
	About          string    `db:"about"`
}

type ProfilePost struct {
	ID            int       `json:"post_id"`
	Title         string    `json:"title"`
	Content       string    `json:"content,omitempty"`
	Image         string    `json:"image_id,omitempty"`
	CreationDate  time.Time `json:"creation_date"`
	PostPrivacy   string    `json:"post_privacy,omitempty"`
	PostLikes     int       `json:"likes_count"`
	CommentsCount int       `json:"comments_count"`
	Likers_ids    []string  `json:"likers_ids"`
	IsLiked       bool      `json:"is_Liked,omitempty"`
}
