package hub

import "fmt"

type RoomExistsError struct {
	roomName string
}

func (ree *RoomExistsError) Error() string {
	return fmt.Sprintf("room by the name %s already exists", ree.roomName)
}

type RoomDoesNotExistsError struct {
	roomName string
}

func (ree *RoomDoesNotExistsError) Error() string {
	return fmt.Sprintf("room by the name %s does not exists", ree.roomName)
}
