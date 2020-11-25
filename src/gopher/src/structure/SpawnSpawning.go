package structure

type SpawnSpawning struct {
	//Directions []int @todo
	Name          string `js:"name"`
	NeedTime      int    `js:"needTime"`
	RemainingTime int    `js:"remainingTime`
	// Spawn         *StructureSpawn
}

// func (spawning *StructureSpawnSpawning) LoadStructureSpawnSpawning(obj js.Value) {
// 	spawning.Name = MarshalString(obj.Get("name"), "")
// 	spawning.NeedTime = MarshalInt(obj.Get("needTime"), 0)
// 	spawning.RemainingTime = MarshalInt(obj.Get("remainingTime"), 0)
// 	if HasKeyOfType(obj, "spawn", js.TypeObject) {
// 		spawning.Spawn.LoadStructureSpawn(obj.Get("spawn"))
// 	}
// }
