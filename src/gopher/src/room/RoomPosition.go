package room

import "github.com/gopherjs/gopherjs/js"

type RoomPosition struct {
	*js.Object
	X        int    `js"x"`
	Y        int    `js:"y"`
	RoomName string `js:"roomName"`
}
