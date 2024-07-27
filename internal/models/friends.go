package models

import "time"

type Friends struct {
	UserName        string          `json:"user_name"`
	Followers       []string        `json:"followers"`
	Following       []string        `json:"following"`
	Friend_requests []FriendRequest `json:"friend_requests"`
	Explore         []string        `json:"explore"` // all the users that are not friends
}

type FriendRequest struct {
	UserName      string    `json:"requester"`
	Creation_date time.Time `json:"creation_date"`
}
