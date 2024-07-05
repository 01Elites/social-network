package models

type Request struct {
	ID       int    `json:"request_id,omitempty"`
	Sender   string `json:"sender,omitempty"`   // user Name
	Receiver string `json:"receiver,omitempty"` // user Name
	Status   string `json:"status,omitempty"`
}

type Response struct {
	ID         int    `json:"response_id,omitempty"`
	FollowerID string `json:"follower_id,omitempty"`
	FolloweeID string `json:"followee_id,omitempty"`
	Status     string `json:"status,omitempty"`
}
