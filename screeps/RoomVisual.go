package screeps

import "syscall/js"

type RoomVisual struct {
	RoomName string
}

func (visual *RoomVisual) LoadRoomVisual(obj js.Value) {
	visual.RoomName = MarshalString(obj.Get("roomName"), "")
}
