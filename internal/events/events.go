package events

type Post struct {
	ID            int    `json:"post_id"`
	User          User   `json:"user"`
	Title         string `json:"title"`
	Content       string `json:"content,omitempty"`
	Category      string `json:"category"`
	CreationDate  string `json:"creation_date"`
	PostLikes     int    `json:"post_likes"`
	IsLiked       bool   `json:"isLiked"`
	CommentsCount int    `json:"comments_count"`
	ViewsCount    int    `json:"views_count"` 
}

type Create_Post struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Privacy  string `json:"post_privacy"`
	GroupID  int    `json:"group_id"`
}

type PostFeed struct {
	ID            int    `json:"post_id"`
	User          User   `json:"user"`
	Title         string `json:"title"`
	Content       string `json:"content,omitempty"`
	Image         string `json:"image"`
	GroupID       int    `json:"group_id"`
	CreationDate  string `json:"creation_date"`
}
type User struct {
	ID           string    `json:"id"`
	UserName     string `json:"username"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Gender       string `json:"gender"`
	// ViewsCount   int    `json:"views_count"`
	// PostsCount   int    `json:"posts_count"`
	// RepliesCount int    `json:"replies_count"`
}