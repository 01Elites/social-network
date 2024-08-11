package models

import "time"

type Request struct {
	ID             int       `json:"request_id,omitempty"`
	Sender         string    `json:"sender,omitempty"`   // user Name
	Receiver       string    `json:"receiver,omitempty"` // user Name
	Status         string    `json:"status,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	NotificationID int       `json:"notification_id,omitempty"`
}

type Response struct {
	ID             int    `json:"response_id,omitempty"`
	Follower       string `json:"follower,omitempty"` // user Name
	Followee       string `json:"followee,omitempty"` // user Name
	Status         string `json:"status,omitempty"`
	NotificationID int    `json:"notification_id,omitempty"`
}

// type Notification struct {
// 	ID         string `json:"id,omitempty"`
// 	UserID     string `json:"user_id,omitempty"`
// 	Type       string `json:"type,omitempty"`
// 	RelatedID  string `json:"relatedID,omitempty"`
// 	Status     string `json:"status,omitempty"`
// }
