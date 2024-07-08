package models

import "time"

type GroupFeed struct {
	ID       int      `json:"id,omitempty"`
	Members  []string `json:"members,omitempty"`
	Posts    []Post   `json:"posts,omitempty"`
	IsMember bool     `json:"ismember,omitempty"`
}

type CreateGroup struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type GroupAction struct {
	ID         int    `json:"id,omitempty"`
	GroupID    int    `json:"group_id,omitempty"`
	Username   string `json:"receiver,omitempty"`
	ReceiverID string `json:"receiver_id,omitempty"`
}
type GroupResponse struct {
	ID             int    `json:"invite_id,omitempty"`
	GroupID        int    `json:"group_id,omitempty"`
	RequesterID    string `json:"requester_id,omitempty"`
	Status         string `json:"response,omitempty"`
	NotificationID int    `json:"notification_id,omitempty"`
}

type CreateEvent struct {
	GroupID     int       `json:"group_id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	EventTime   time.Time `json:"event_date,omitempty"`
}

type EventResp struct {
	EventID int `json:"event_id,omitempty"`
	OptionID int `json:"option_id,omitempty"`
}
