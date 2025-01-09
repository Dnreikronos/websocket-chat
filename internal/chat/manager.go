package chat

type Manager struct {
	Clients    map[string]*Client
	Broadcast  chan Message
	Register   chan *Client
	Unregister chan *Client
}

func NewManager() *Manager {
	return &Manager{
		Clients:    make(map[string]*Client),
		Broadcast:  make(chan Message),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (m *Manager) Run() {
	for {
		select {
		case client := <-m.Register:
			m.Clients[client.ID] = client
		case client := <-m.Unregister:
			if _, ok := m.Clients[client.ID]; ok {
				delete(m.Clients, client.ID)
				close(client.Send)
			}
		case msg := <-m.Broadcast:
			if client, ok := m.Clients[msg.ReceiverID]; ok {
				client.Send <- msg
			}
		}
	}
}
