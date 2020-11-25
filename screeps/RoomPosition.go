package screeps

import "syscall/js"

type RoomPosition struct {
	X        int
	Y        int
	RoomName string
}

func (pos *RoomPosition) LoadRoomPosition(obj js.Value) {
	pos.X = MarshalInt(obj.Get("x"), 0)
	pos.Y = MarshalInt(obj.Get("y"), 0)
	pos.RoomName = MarshalString(obj.Get("roomName"), "")
}
