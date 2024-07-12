package types

import "github.com/gorilla/websocket"

type Section struct {
	Name  string `json:"name"`
	Users []UserDetails `json:"users"`
}

type User struct {
	ID       string          `json:"-"`
	Username string          `json:"username"`
	Conn     *websocket.Conn `json:"-"`
	ChatOpened string          `json:"-"`

}

type UserDetails struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	State     string `json:"state,omitempty"`
}
