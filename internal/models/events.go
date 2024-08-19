package models

import "time"

type CreateEvent struct {
	GroupID     int       `json:"group_id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	Options     []string  `json:"options,omitempty"`
	EventTime   time.Time `json:"event_date,omitempty"`
}

type EventResp struct {
	EventID  int `json:"event_id,omitempty"`
	OptionID int `json:"option_id,omitempty"`
}

type Event struct {
	ID          int             `json:"id,omitempty"`
	Title       string          `json:"title,omitempty"`
	Description string          `json:"description,omitempty"`
	Creator     PostFeedProfile `json:"creator,omitempty"`
	Options     []Options       `json:"options,omitempty"`
	EventTime   time.Time       `json:"event_time,omitempty"`
}

type Options struct {
	ID        int      `json:"option_id,omitempty"`
	Name      string   `json:"option_name,omitempty"`
	Usernames []string `json:"usernames"`
	FullNames []string `json:"fullnames"`
}

type CancelEvent struct {
	EventID int `json:"event_id,omitempty"`
}

type MyEvents struct {
	PendingEvents  []Event `json:"pending"`
	UpcomingEvents []Event `json:"upcoming"`
	PastEvents     []Event `json:"past"`
}
