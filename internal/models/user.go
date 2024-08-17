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

// GiteaUser struct defines the structure for storing user information from Gitea
type GiteaUser = struct {
	Login     string `json:"login"`
	FullName  string `json:"full_name"`
	Email     string `json:"email"`
	AvatarURL string `json:"avatar_url"`
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
	UserID         string    `db:"user_id" json:"user_id,omitempty"`
	Username       string    `db:"user_name" json:"user_name,omitempty"`
	NickName       string    `db:"nick_name" json:"nick_name,omitempty"`
	FirstName      string    `db:"first_name" json:"first_name"`
	LastName       string    `db:"last_name" json:"last_name"`
	Gender         string    `db:"gender" json:"gender,omitempty"`
	DateOfBirth    time.Time `db:"date_of_birth" json:"date_of_birth,omitempty"`
	ProfilePrivacy string    `db:"privacy" json:"post_privacy,omitempty"`
	Avatar         string    `db:"image" json:"avatar,omitempty"`
	About          string    `db:"about" json:"about,omitempty"`
}

type ProfilePost struct {
	ID            int       `json:"post_id"`
	Title         string    `json:"title"`
	Content       string    `json:"content,omitempty"`
	Image         string    `json:"image,omitempty"`
	CreationDate  time.Time `json:"creation_date"`
	PostPrivacy   string    `json:"post_privacy,omitempty"`
	PostLikes     int       `json:"likes_count"`
	CommentsCount int       `json:"comments_count"`
	Likers_ids    []string  `json:"likers_ids,omitempty"`
	IsLiked       bool      `json:"is_Liked,omitempty"`
}

type PostFeedProfile struct {
	UserID         string `db:"user_id" json:"user_id,omitempty"`
	UserName       string `db:"user_name" json:"user_name,omitempty"`
	FirstName      string `db:"first_name" json:"first_name"`
	LastName       string `db:"last_name" json:"last_name"`
	ProfilePrivacy string `db:"privacy" json:"post_privacy,omitempty"`
	Avatar         string `db:"image" json:"avatar,omitempty"`
}

var Code string
var Token string
