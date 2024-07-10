package types

import "github.com/gorilla/websocket"

type Section struct {
	Name  string `json:"name"`
	Users []User `json:"users"`
}

type User struct {
	ID       string          `json:"-"`
	Username string          `json:"username"`
	State    string          `json:"state"`
	Conn     *websocket.Conn `json:"-"`
}
