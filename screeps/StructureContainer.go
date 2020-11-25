package screeps

import "syscall/js"

type StructureContainer struct {
	Structure

	Store        Store
	TicksToDecay int
}

func (structure *StructureContainer) LoadStructureContainer(obj js.Value) {
	structure.LoadStructure(obj)

	if HasKeyOfType(obj, "store", js.TypeObject) {
		structure.Store.LoadStore(obj.Get("store"))
	}
	structure.TicksToDecay = MarshalInt(obj.Get("ticksToDecay"), 0)
}
