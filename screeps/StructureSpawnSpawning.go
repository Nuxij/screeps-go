package screeps

import "syscall/js"

type StructureSpawnSpawning struct {
	//Directions []int @todo
	Name          string
	NeedTime      int
	RemainingTime int
	Spawn         *StructureSpawn
}

func (spawning *StructureSpawnSpawning) LoadStructureSpawnSpawning(obj js.Value) {
	spawning.Name = MarshalString(obj.Get("name"), "")
	spawning.NeedTime = MarshalInt(obj.Get("needTime"), 0)
	spawning.RemainingTime = MarshalInt(obj.Get("remainingTime"), 0)
	if HasKeyOfType(obj, "spawn", js.TypeObject) {
		spawning.Spawn.LoadStructureSpawn(obj.Get("spawn"))
	}
}
