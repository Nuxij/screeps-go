package screeps

import "syscall/js"

type StructureTower struct {
	OwnedStructure

	Store Store
}

func (structure *StructureTower) LoadStructureTower(obj js.Value) {
	structure.LoadOwnedStructure(obj)

	if HasKeyOfType(obj, "store", js.TypeObject) {
		structure.Store.LoadStore(obj.Get("store"))
	}
}
