package screeps

import "syscall/js"

type Structure struct {
	RoomObject

	Hits          int
	HitsMax       int
	ID            string
	StructureType string // @todo could be constant
}

func (structure *Structure) LoadStructure(obj js.Value) {
	structure.LoadRoomObject(obj)

	structure.Hits = MarshalInt(obj.Get("hits"), 0)
	structure.HitsMax = MarshalInt(obj.Get("hitsMax"), 0)
	structure.ID = MarshalString(obj.Get("id"), "")
	structure.StructureType = MarshalString(obj.Get("structureType"), "")
}
