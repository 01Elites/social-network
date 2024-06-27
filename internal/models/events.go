package models

import "time"

type Post struct {
	ID            int         `json:"post_id"`
	User          UserProfile `json:"user"`
	Title         string      `json:"title"`
	Content       string      `json:"content,omitempty"`
	Image         string      `json:"image"`
	CreationDate  time.Time   `json:"creation_date"`
	PostPrivacy   string      `json:"post_privacy"`
	PostLikes     int         `json:"post_likes"`
	GroupID       int         `json:"group_id"`
	IsLiked       bool        `json:"isLiked"`
	CommentsCount int         `json:"comments_count"`
	Likers_ids    []string    `json:"likers_ids"`
}

type Create_Post struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Privacy string   `json:"privacy"`
	Image   string   `json:"image_id"`
	GroupID int      `json:"group_id"`
	UserIDs []string `json:"user_ids"`
}

// type PostFeed struct {
// 	User          User      `json:"user"`
// 	Title         string    `json:"title"`
// 	Content       string    `json:"content,omitempty"`
// 	Image         string    `json:"image"`
// 	GroupID       int       `json:"group_id"`
// 	CreationDate  time.Time `json:"creation_date"`
// 	CommentsCount int       `json:"comments_count"`
// 	Likers_ids    []string  `json:"likers_ids"`
// }

type Create_Comment struct {
	ParentID int    `json:"parent_id"`
	Content  string `json:"content"`
}

type Comment struct {
	ID           int    `json:"comment_id"`
	User         User   `json:"user"`
	Content      string `json:"content"`
	CreationDate string `json:"creation_date"`
	CommentLikes int    `json:"comment_likes"`
	IsLiked      bool   `json:"isLiked"`
}

type Create_Group struct {
	Title       string
	Description string
}

type Error struct {
	Reason string `json:"reason"`
}
