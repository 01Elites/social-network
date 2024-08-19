package models

type GroupFeed struct {
	Title       string            `json:"title,omitempty"`
	Description string            `json:"description,omitempty"`
	ID          int               `json:"id,omitempty"`
	Members     []UserLiteInfo    `json:"members,omitempty"`
	Events      []Event           `json:"events,omitempty"`
	IsMember    bool              `json:"ismember"`
	IsCreator   bool              `json:"iscreator"`
	RequestMade bool              `json:"request_made"`
	InvitedBy   Requester         `json:"invited_by,omitempty"`
	Creator     PostFeedProfile   `json:"creator,omitempty"`
	Requesters  []Requester       `json:"requesters,omitempty"`
	Explore     []PostFeedProfile `json:"explore,omitempty"`
}

type CreateGroup struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}

type Requester struct {
	User         PostFeedProfile `json:"user,omitempty"`
	CreationDate string          `json:"creation_date"`
}

type GroupAction struct {
	ID int `json:"id,omitempty"`
	// GroupIDstr string `json:"group_id,omitempty"`
	GroupID    int    `json:"group_id,omitempty"`
	Username   string `json:"receiver,omitempty"`
	ReceiverID string `json:"receiver_id,omitempty"`
}
type GroupResponse struct {
	ID          int    `json:"invite_id,omitempty"`
	GroupID     int    `json:"group_id,omitempty"`
	Requester   string `json:"requester,omitempty"`
	RequesterID string
	Status      string `json:"response,omitempty"`
}

type Groups struct {
	Owned   []GroupFeed `json:"owned,omitempty"`
	Joined  []GroupFeed `json:"joined,omitempty"`
	Explore []GroupFeed `json:"explore,omitempty"`
}
