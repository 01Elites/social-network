package models

import "time"

type GroupFeed struct {
	Title       string          `json:"title,omitempty"`
	Description string          `json:"description,omitempty"`
	ID          int             `json:"id,omitempty"`
	Members     []string        `json:"members,omitempty"`
	Events      []Event         `json:"events,omitempty"`
	IsMember    bool            `json:"ismember"`
	IsCreator   bool            `json:"iscreator"`
	RequestMade bool            `json:"request_made"`
	InvitedBy		string					`json:"invited_by,omitempty"`
	Creator     PostFeedProfile `json:"creator,omitempty"`
	Requesters  []Requester     `json:"requesters,omitempty"`
}

type CreateGroup struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type Requester struct {
	User         PostFeedProfile `json:"user,omitempty"`
	CreationDate time.Time       `json:"creation_date"`
}

type GroupAction struct {
	ID         int    `json:"id,omitempty"`
	GroupID    int    `json:"group_id,omitempty"`
	Username   string `json:"receiver,omitempty"`
	ReceiverID string `json:"receiver_id,omitempty"`
}
type GroupResponse struct {
	ID          int    `json:"invite_id,omitempty"`
	GroupIdstr  string `json:"group_id,omitempty"`
	GroupID     int    
	Requester   string `json:"requester,omitempty"`
	RequesterID string
	Status      string `json:"response,omitempty"`
}

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
	ID             int       `json:"id,omitempty"`
	Title          string    `json:"title,omitempty"`
	Description    string    `json:"description,omitempty"`
	Options        []string  `json:"options,omitempty"`
	EventTime      time.Time `json:"event_time,omitempty"`
	RespondedUsers []string  `json:"responded_users,omitempty"`
}

type CancelEvent struct {
	EventID int `json:"event_id,omitempty"`
}
