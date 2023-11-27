package hub

type Room struct {
	ID        string
	Name      string
	Clients   map[string]*Client
	Broadcast chan *Message
}

func NewRoom(id string, name string) *Room {
	room := &Room{
		ID:        id,
		Name:      name,
		Clients:   make(map[string]*Client),
		Broadcast: make(chan *Message, 10),
	}

	go room.eventLoop()

	return room
}

func (r *Room) AddClient(client *Client) {
	client.BelongTo[r.ID] = r
	r.Clients[client.ID] = client
}

func (r *Room) RemoveClient(client *Client) {
	delete(r.Clients, client.ID)
}

func (r *Room) eventLoop() {
	for {
		select {
		case msg := <-r.Broadcast:
			for _, client := range r.Clients {
				client.Send(msg)
			}
		}
	}
}
