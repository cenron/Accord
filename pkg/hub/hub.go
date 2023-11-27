package hub

type Hub struct {
	Rooms map[string]*Room
}

func NewHub() *Hub {
	return &Hub{
		Rooms: make(map[string]*Room),
	}
}

func (h *Hub) CreateRoom(roomName string) (*Room, error) {
	if h.RoomNameExists(roomName) {
		return nil, &RoomExistsError{roomName: roomName}
	}

	// Create a unique room id.
	roomID := "1234" //uuid.New().String()

	// Create our new room
	room := NewRoom(roomID, roomName)

	// Add our room to the hub.
	h.Rooms[roomID] = room

	return room, nil
}

func (h *Hub) GetRoomByName(roomName string) (*Room, error) {
	for _, room := range h.Rooms {
		if roomName == room.Name {
			return room, nil
		}
	}

	return nil, &RoomDoesNotExistsError{roomName: roomName}
}

func (h *Hub) RoomIdExists(roomID string) bool {
	if _, ok := h.Rooms[roomID]; ok {
		return true
	}

	return false
}

func (h *Hub) RoomNameExists(roomName string) bool {
	for _, room := range h.Rooms {
		if roomName == room.Name {
			return true
		}
	}

	return false
}
