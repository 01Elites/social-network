package models

import "time"

type Friends struct {
	UserName        string          `json:"user_name,omitempty"`
	Followers       []UserLiteInfo  `json:"followers"`
	Following       []UserLiteInfo  `json:"following"`
	Friend_requests []FriendRequest `json:"friend_requests"`
	Explore         []UserLiteInfo  `json:"explore"` // all the users that are not friends
}

type FriendRequest struct {
	UserName      string       `json:"requester"`
	Creation_date time.Time    `json:"creation_date"`
	UserInfo      UserLiteInfo `json:"user_info,omitempty"`
}

type UserLiteInfo struct {
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
}
