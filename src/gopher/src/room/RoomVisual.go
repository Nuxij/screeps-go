package room

import "github.com/gopherjs/gopherjs/js"

type RoomVisual struct {
	*js.Object
	RoomName string `js:"roomName"`
}
