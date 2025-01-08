package chat

import "github.com/gorilla/websocket"

type Client struct {
	ID     string
	Socket *websocket.Conn
	Send   chan Message
}
