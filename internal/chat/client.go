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

	for {
		var msg Message
		if err := c.Socket.ReadJSON(&msg); err != nil {
			break
		}
		manager.Broadcast <- msg
	}
}

func (c *Client) Write() {
	defer c.Socket.Close()
	for msg := range c.Send() {
		c.Socket.WriteJSON(msg)
	}
}
