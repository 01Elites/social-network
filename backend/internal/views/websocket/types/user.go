package types

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Section struct {
	Name  string        `json:"name"`
	Users []UserDetails `json:"users"`
}

type User struct {
	ID                string                   `json:"-"`
	Username          string                   `json:"username"`
	Conns             map[*websocket.Conn]bool `json:"-"`
	Mutex             *sync.Mutex              `json:"-"`
	ChatOpened        string                   `json:"-"`
	ChatOpenedIsGroup bool                     `json:"-"`
}

type UserDetails struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Username  string `json:"user_name,omitempty"`
	State     string `json:"state,omitempty"`
	Avatar    string `json:"avatar,omitempty"`
}

type UserList struct {
	Type     string      `json:"type"`
	Metadata interface{} `json:"metadata"`
}
