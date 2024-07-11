package types

type Event struct {
	Type    string      `json:"event"`
	Id      string      `json:"id"`
	Payload interface{} `json:"payload"`
}

// notification payload struct
type Notification struct {
	Type     string      `json:"type"`
	Message  string      `json:"message"`
	Metadata interface{} `json:"metadata"`
}

// type FollowRequestMetadata struct {
// 	UserDetails UserDetails `json:"requester"`
// }

// type comment notification metadata struct
type CommentMetadata struct {
	PostID    string    `json:"post_id"`
	Commenter Commenter `json:"commenter"`
}

type Commenter struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	ID        string `json:"id"`
	ImageID   string `json:"image_id"`
}
