package models

type GroupFeed struct {
	ID          int
	Members     []User
	Posts       []Post 
	IsMember    bool
}