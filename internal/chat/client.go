package chat

import "github.com/gorilla/websocket"

type Client struct {
	ID     string
	Socket *websocket.Conn
	Send   chan Message
}

func (c *Client) Read(manager *Manager) {
	defer func() {
		manager.Unregister <- c
		c.Socket.Close()
	}()
}
