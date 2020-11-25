package screeps

import "syscall/js"

type StructureWall struct {
	OwnedStructure
}

func (structure *StructureWall) LoadStructureWall(obj js.Value) {
	structure.LoadOwnedStructure(obj)
}
