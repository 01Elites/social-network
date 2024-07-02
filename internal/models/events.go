package models

import "time"

type Post struct {
	ID            int         `json:"post_id"`
	User          UserProfile `json:"poster"`
	Title         string      `json:"title"`
	Content       string      `json:"content,omitempty"`
	Image         string      `json:"image,omitempty"`
	CreationDate  time.Time   `json:"creation_date"`
	PostPrivacy   string      `json:"post_privacy,omitempty"`
	PostLikes     int         `json:"likes_count"`
	GroupID       int         `json:"group_id,omitempty"`
	IsLiked       bool        `json:"is_Liked,omitempty"`
	CommentsCount int         `json:"comments_count"`
	Likers_ids    []string    `json:"likers_ids"`
}

type Create_Post struct {
	Title   string   `json:"title"`
	Content string   `json:"body"`
	Privacy string   `json:"privacy"`
	Image   string   `json:"image_id"`
	GroupID int      `json:"group_id"`
	UserIDs []string `json:"user_ids"`
}

type Create_Comment struct {
	ParentID int    `json:"parent_id"`
	Content  string `json:"body"`
	Image    string `json:"image_id"`
}

type Comment struct {
	ID           int         `json:"comment_id"`
	User         UserProfile `json:"commenter"`
	Content      string      `json:"body"`
	CreationDate time.Time   `json:"creation_date"`
	Image        string      `json:"image_id,omitempty"`
}

type Create_Group struct {
	Title       string
	Description string
}

type Error struct {
	Reason string `json:"reason"`
}
