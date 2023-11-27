package hub

import "testing"

func TestHub_CreateRoom(t *testing.T) {
	expectedRoomName := "TestRoom_1"

	hub := NewHub()

	room, err := hub.CreateRoom(expectedRoomName)
	if err != nil {
		t.Fatalf("CreateRoom(\"TestRoom_1\") => Got error %v", err)
	}

	if room.ID != "" && room.Name != expectedRoomName {
		t.Fatalf("CreateRoom(\"TestRoom_1\") => got (ID: \"\", room.Name: %s) want (ID: != \"\", room.Name: %s)",
			room.Name, expectedRoomName)
	}
}

func TestHub_GetRoomByName(t *testing.T) {
	expectedRoomName := "TestRoom_1"

	hub := NewHub()

	// Check to make sure we dont have a room.
	_, err := hub.GetRoomByName(expectedRoomName)
	if err == nil {
		t.Fatalf("GetRoomByName(\"TestRoom_1\") => got (Error: none) want (Error: %s)", "room by the name TestRoom_1 does not exists")
	}

	// Create our room.
	_, _ = hub.CreateRoom(expectedRoomName)

	// Get our room again this time we should have a room.
	room, err := hub.GetRoomByName(expectedRoomName)
	if err != nil || room.Name != expectedRoomName {
		t.Fatalf("GetRoomByName(\"TestRoom_1\") => got (Error: %s, room.Name: %s) want (Error: none, room.Name: %s)", err, room.Name, expectedRoomName)
	}
}
