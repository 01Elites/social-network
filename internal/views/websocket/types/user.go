package types

import "github.com/gorilla/websocket"

type Section struct {
	User         User `json:"user"`
	Notification bool `json:"notification"`
}

type User struct {
	Username string          `json:"username"`
	State    string          `json:"state"`
	Conn     *websocket.Conn `json:"-"`
}
