package hub

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
)

type Message struct {
	RoomID   string `json:"roomID"`
	ClientID string `json:"clientID"`
	Content  string `json:"content"`
}

type Client struct {
	ID       string
	Username string
	Socket   *websocket.Conn
	BelongTo map[string]*Room
}

func NewClient(id string, username string, socket *websocket.Conn) *Client {
	client := &Client{
		ID:       id,
		Username: username,
		Socket:   socket,
		BelongTo: make(map[string]*Room),
	}

	go client.readMessage()

	return client
}

func (cl *Client) Send(message *Message) {
	_ = cl.Socket.WriteJSON(message)
}

func (cl *Client) readMessage() {
	defer func() {
		_ = cl.Socket.Close()
	}()

	for {
		_, msg, err := cl.Socket.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("connection closed: %v", err)
			}
			break
		}

		var message Message
		if err := json.Unmarshal(msg, &message); err != nil {
			log.Printf("could not parse message")
			break
		}

		for roomID, room := range cl.BelongTo {
			if roomID != message.RoomID {
				continue
			}

			room.Broadcast <- &message
		}
	}
}
