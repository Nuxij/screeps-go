package game

import (
	"game/memory"
	"room"
)

type Creep struct {
	*room.RoomObject
	Name string `js:"name"`
	// Body    []*CreepBodyPart `js:"body"`
	Fatigue int                 `js:"fatigue"`
	Hits    int                 `js:"hits"`
	HitsMax int                 `js:"hitsMax"`
	ID      string              `js:"id"`
	Memory  *memory.CreepMemory `js:"memory"`
	My      bool                `js:"my"`
	// Owner    *Owner `js:"owner"`
	Saying   string `js:"saying"`
	Spawning bool   `js:"spawning"`
	//Store store @todo
	TicksToLive int `js:"ticksToLive"`
}
