package screeps

import "syscall/js"

type StructureSpawn struct {
	OwnedStructure

	//Memory @todo
	Name     string
	Spawning StructureSpawnSpawning
	Store    Store
}

func (structure *StructureSpawn) LoadStructureSpawn(obj js.Value) {
	structure.LoadOwnedStructure(obj)

	structure.Name = MarshalString(obj.Get("name"), "")
	if HasKeyOfType(obj, "spawning", js.TypeObject) {
		structure.Spawning.LoadStructureSpawnSpawning(obj.Get("spawning"))
	}
	if HasKeyOfType(obj, "store", js.TypeObject) {
		structure.Store.LoadStore(obj.Get("store"))
	}
}
