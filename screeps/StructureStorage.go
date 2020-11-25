package screeps

import "syscall/js"

type StructureStorage struct {
	OwnedStructure

	Store Store
}

func (structure *StructureStorage) LoadStructureStorage(obj js.Value) {
	structure.LoadOwnedStructure(obj)

	if HasKeyOfType(obj, "store", js.TypeObject) {
		structure.Store.LoadStore(obj.Get("store"))
	}
}
