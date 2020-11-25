package screeps

import "syscall/js"

type StructurePowerSpawn struct {
	OwnedStructure

	Store Store
}

func (structure *StructurePowerSpawn) LoadStructurePowerSpawn(obj js.Value) {
	structure.LoadOwnedStructure(obj)

	if HasKeyOfType(obj, "store", js.TypeObject) {
		structure.Store.LoadStore(obj.Get("store"))
	}
}
