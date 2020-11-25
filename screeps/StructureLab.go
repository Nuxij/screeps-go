package screeps

import "syscall/js"

type StructureLab struct {
	OwnedStructure

	Cooldown int
	Store    Store
}

func (structure *StructureLab) LoadStructureLab(obj js.Value) {
	structure.LoadOwnedStructure(obj)

	structure.Cooldown = MarshalInt(obj.Get("cooldown"), 0)
	if HasKeyOfType(obj, "store", js.TypeObject) {
		structure.Store.LoadStore(obj.Get("store"))
	}
}
