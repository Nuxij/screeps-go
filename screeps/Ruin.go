package screeps

import "syscall/js"

type Ruin struct {
	RoomObject

	DestroyTime  int
	ID           string
	Store        Store
	Structure    Structure
	TicksToDecay int
}

func (ruin *Ruin) LoadRuin(obj js.Value) {
	ruin.LoadRoomObject(obj)

	ruin.DestroyTime = MarshalInt(obj.Get("destroyTime"), 0)
	ruin.ID = MarshalString(obj.Get("id"), "")
	if HasKeyOfType(obj, "store", js.TypeObject) {
		ruin.Store.LoadStore(obj.Get("store"))
	}
	if HasKeyOfType(obj, "structure", js.TypeObject) {
		ruin.Structure.LoadStructure(obj.Get("structure"))
	}
	ruin.TicksToDecay = MarshalInt(obj.Get("ticksToDecay"), 0)
}
