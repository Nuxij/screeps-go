package screeps

import "syscall/js"

type StructureRampart struct {
	OwnedStructure

	IsPublic     bool
	TicksToDecay int
}

func (structure *StructureRampart) LoadStructureRampart(obj js.Value) {
	structure.LoadOwnedStructure(obj)

	structure.IsPublic = MarshalBool(obj.Get("isPublic"), false)
	structure.TicksToDecay = MarshalInt(obj.Get("ticksToDecay"), 0)
}
