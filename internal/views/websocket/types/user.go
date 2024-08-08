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
	ID                string          `json:"-"`
	Username          string          `json:"username"`
	Conn              *websocket.Conn `json:"-"`
	Mutex             *sync.Mutex     `json:"-"`
	ChatOpened        string          `json:"-"`
	ChatOpenedIsGroup bool            `json:"-"`
}

type UserDetails struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	State     string `json:"state,omitempty"`
	Avatar		string `json:"avatar,omitempty"`
}
