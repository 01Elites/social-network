package events

type Post struct {
	ID            int    `json:"post_id"`
	User          User   `json:"user"`
	Title         string `json:"title"`
	Content       string `json:"content,omitempty"`
	Category      string `json:"category"`
	CreationDate  string `json:"creation_date"`
	PostPrivacy   string `json:"post_privacy"`
	PostLikes     int    `json:"post_likes"`
	GroupID       int    `json:"group_id"`
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
	Following    map[string]bool `json:"following"`
	Groups       map[int]bool `json:"groups"`
	// ViewsCount   int    `json:"views_count"`
	// PostsCount   int    `json:"posts_count"`
	// RepliesCount int    `json:"replies_count"`
}

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
	Title 			string
	Description string
}