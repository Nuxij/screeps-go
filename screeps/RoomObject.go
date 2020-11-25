package screeps

import "syscall/js"

type RoomObject struct {
	Effects []Effect
	Pos     RoomPosition
	//Room    Room    @todo  we need this to work but we get illegal cycle in declaration
}

func (roomObj *RoomObject) LoadRoomObject(obj js.Value) {
	roomObj.Effects = MarshalEffects(obj.Get("effects"))
	if HasKeyOfType(obj, "pos", js.TypeObject) {
		roomObj.Pos.LoadRoomPosition(obj.Get("pos"))
	}
	if HasKeyOfType(obj, "room", js.TypeObject) {
		//roomObj.Room.LoadRoom(obj.Get("room"))
	}
}
