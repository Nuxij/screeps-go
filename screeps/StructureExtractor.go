package screeps

import "syscall/js"

type StructureExtractor struct {
	OwnedStructure

	Cooldown int
}

func (structure *StructureExtractor) LoadStructureExtractor(obj js.Value) {
	structure.LoadOwnedStructure(obj)

	structure.Cooldown = MarshalInt(obj.Get("cooldown"), 0)
}
