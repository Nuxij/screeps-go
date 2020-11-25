package screeps

import "syscall/js"

type CreepBodyPart struct {
	Boost string // @todo could be a real constant
	Type  string
	Hits  int
}

func (body *CreepBodyPart) LoadCreepBodyPart(obj js.Value) {
	body.Boost = MarshalString(obj.Get("boost"), "")
	body.Type = MarshalString(obj.Get("type"), "")
	body.Hits = MarshalInt(obj.Get("hits"), 0)
}
