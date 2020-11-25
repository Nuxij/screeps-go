package screeps

import "syscall/js"

type StructureObserver struct {
	OwnedStructure
}

func (structure *StructureObserver) LoadStructureObserver(obj js.Value) {
	structure.LoadOwnedStructure(obj)
}
