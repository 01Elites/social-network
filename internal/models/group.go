package models

type GroupFeed struct {
	ID       int    `json:"id,omitempty"`
	Members  []User `json:"members,omitempty"`
	Posts    []Post `json:"posts,omitempty"`
	IsMember bool   `json:"ismember,omitempty"`
}


type Create_Group struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
}