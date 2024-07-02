package models

type Request struct {
	ID          int    `json:"request_id,omitempty"`
	SenderID    string `json:"sender_id,omitempty"`
	ReceiverID  string `json:"receiver_id,omitempty"`
	Status		  string `json:"status,omitempty"`
}

type Response struct {
	ID    			int    `json:"response_id,omitempty"`
	FollowerID  string `json:"follower_id,omitempty"`
	FolloweeID  string `json:"followee_id,omitempty"`
	Status		  string `json:"status,omitempty"`
}