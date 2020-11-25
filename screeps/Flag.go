package screeps

import "syscall/js"

type Flag struct {
	RoomObject

	Color int // @todo could be a constant
	//Memory @todo missing
	Name           string
	SecondaryColor int // @todo could be a constant
}

func (flag *Flag) LoadFlag(obj js.Value) {
	flag.LoadRoomObject(obj)

	flag.Color = MarshalInt(obj.Get("color"), 0)
	flag.Name = MarshalString(obj.Get("name"), "")
	flag.SecondaryColor = MarshalInt(obj.Get("secondaryColor"), 0)
}
