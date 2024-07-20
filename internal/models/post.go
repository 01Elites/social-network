package models

import "time"

type Post struct {
	ID               int             `json:"post_id"`
	User             PostFeedProfile `json:"poster,omitempty"`
	Title            string          `json:"title"`
	Content          string          `json:"content,omitempty"`
	Image            string          `json:"image,omitempty"`
	CreationDate     time.Time       `json:"creation_date"`
	PostPrivacy      string          `json:"post_privacy,omitempty"`
	PostLikes        int             `json:"likes_count"`
	CommentsCount    int             `json:"comments_count"`
	Likers_Usernames []string        `json:"likers_usernames"`
	IsLiked          bool            `json:"is_liked,omitempty"`
	GroupID          int             `json:"group_id,omitempty"`
}

type Create_Post struct {
	Title     string   `json:"title"`
	Content   string   `json:"body"`
	Privacy   string   `json:"privacy"`
	Image     string   `json:"image"`
	GroupID   int      `json:"group_id"`
	UserNames []string `json:"usernames"`
}

type Create_Comment struct {
	ParentID int    `json:"parent_id"`
	Content  string `json:"body"`
	Image    string `json:"image"`
}

type Comment struct {
	ID           int             `json:"comment_id"`
	User         PostFeedProfile `json:"commenter"`
	Content      string          `json:"body"`
	CreationDate time.Time       `json:"creation_date"`
	Image        string          `json:"image,omitempty"`
}

type Error struct {
	Reason string `json:"reason"`
}

type ID struct {
	ID int `json:"id"`
}
