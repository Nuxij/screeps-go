package screeps

import "syscall/js"

type StructureExtension struct {
	OwnedStructure

	Store Store
}

func (structure *StructureExtension) LoadStructureExtension(obj js.Value) {
	structure.LoadOwnedStructure(obj)

	if HasKeyOfType(obj, "store", js.TypeObject) {
		structure.Store.LoadStore(obj.Get("store"))
	}
}
