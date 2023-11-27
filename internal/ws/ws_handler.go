package ws

import (
	"accord/pkg/db"
	"accord/pkg/hub"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:    1024,
	WriteBufferSize:   1024,
	EnableCompression: true,
}

type WsHandler struct {
	store *db.MongoStore
	hub   *hub.Hub
	room  *hub.Room
}

func NewWsHandler(store *db.MongoStore) *WsHandler {
	clientHub := hub.NewHub()
	room, _ := clientHub.CreateRoom("Test Room")

	return &WsHandler{
		store: store,
		hub:   clientHub,
		room:  room,
	}
}

func (ws *WsHandler) JoinRoom(w http.ResponseWriter, r *http.Request) {

	var query = r.URL.Query()
	clientID := query["clientID"][0]

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client := hub.NewClient(clientID, clientID, conn)

	ws.room.AddClient(client)

	/*	m := &hub.Message{
			RoomID:   "1234",
			ClientID: "5678",
			Content:  "Hello world",
		}

		room.Broadcast <- m*/
	/*	defer func() {
			_ = conn.Close()
		}()

		m := &hub.Message{
			RoomID:   "1234",
			ClientID: "5678",
			Content:  "Hello world",
		}

		_ = conn.WriteJSON(m)*/

}
