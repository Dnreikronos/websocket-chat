package chat

type Manager struct {
	Clients    map[string]*Client
	Broadcast  chan Message
	Register   chan *Client
	Unregister chan *Client
}
