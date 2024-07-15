package types

type Event struct {
	Type    string      `json:"event"`
	Id      string      `json:"id,omitempty"`
	ToUser  string      `json:"username,omitempty"`
	Payload interface{} `json:"payload"`
}

// notification payload struct
type Notification struct {
	Type     string      `json:"type"`
	Message  string      `json:"message"`
	ToUser   string      `json:"-"`
	ID       int         `json:"notification_id"`
	Read     bool        `json:"read"`
	Metadata interface{} `json:"metadata"`
}

type FollowRequestMetadata struct {
	UserDetails UserDetails `json:"requester"`
}

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

type Chat struct {
	ID        int    `json:"id,omitempty"`
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Message   string `json:"message,omitempty"`
	Date      string `json:"date,omitempty"`
	Read      bool   `json:"read,omitempty"`
}

type Typing struct {
	Recipient string `json:"recipient"`
}
