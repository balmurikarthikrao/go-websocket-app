package models

import (
	"github.com/gorilla/websocket"
)

type User struct {
	ID         string
	Connection *websocket.Conn
}
