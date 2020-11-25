package structure

import (
	"room"
)

type Spawn struct {
	*Structure

	//Memory @todo
	ID            string `js:"id"`
	Progress      int    `js:"progress"`
	ProgressTotal int    `js:"progressTotal"`
	// Spawning      SpawnSpawning `js:"spawning"`
	// Store         Store         `js:"Store"`
	Room *room.Room `js:"room"`
}

// func (s *Spawn) Room() *room.Room {
// 	return
// }

// func (structure *StructureSpawn) LoadStructureSpawn(obj js.Value) {
// 	structure.LoadOwnedStructure(obj)

// 	structure.Name = MarshalString(obj.Get("name"), "")
// 	if HasKeyOfType(obj, "spawning", js.TypeObject) {
// 		structure.Spawning.LoadStructureSpawnSpawning(obj.Get("spawning"))
// 	}
// 	if HasKeyOfType(obj, "store", js.TypeObject) {
// 		structure.Store.LoadStore(obj.Get("store"))
// 	}
// }

// func (structure *StructureSpawn) SpawnCreep(parts []string, name string) int {
// 	g := js.Global.Get("Game")
// 	// fmt.Printf("%+v", g)
// 	s := g.Get("spawns").Get(structure.Name)
// 	// fmt.Printf("%+v", s)

// 	c := s.Call("spawnCreep", parts, name)
// 	// fmt.Printf("%+v", c)
// 	return c.Int()
// }
