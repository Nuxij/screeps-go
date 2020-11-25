package screeps

import "syscall/js"

type StructureRoad struct {
	OwnedStructure

	TicksToDecay int
}

func (structure *StructureRoad) LoadStructureRoad(obj js.Value) {
	structure.LoadOwnedStructure(obj)

	structure.TicksToDecay = MarshalInt(obj.Get("ticksToDecay"), 0)
}
