package screeps

import "syscall/js"

type StructureFactory struct {
	OwnedStructure

	Cooldown int
	Level    int
	Store    Store
}

func (structure *StructureFactory) LoadStructureFactory(obj js.Value) {
	structure.LoadOwnedStructure(obj)

	structure.Cooldown = MarshalInt(obj.Get("cooldown"), 0)
	structure.Level = MarshalInt(obj.Get("level"), 0)
	if HasKeyOfType(obj, "store", js.TypeObject) {
		structure.Store.LoadStore(obj.Get("store"))
	}
}
