package creep

import (
	"room"
)

type Creep struct {
	*room.RoomObject
	Name string `js:"name"`
	// Body    []*CreepBodyPart `js:"body"`
	Fatigue int    `js:"fatigue"`
	Hits    int    `js:"hits"`
	HitsMax int    `js:"hitsMax"`
	ID      string `js:"id"`
	My      bool   `js:"my"`
	// Owner    *Owner `js:"owner"`
	Saying   string `js:"saying"`
	Spawning bool   `js:"spawning"`
	//Store store @todo
	TicksToLive int `js:"ticksToLive"`
}

func (c *Creep) Harvest(s *room.RoomObject) int {
	return c.Call("harvest", s.Object).Int()
}

func (c *Creep) MoveTo(pos *room.RoomPosition) int {
	return c.Call("moveTo", pos).Int()
}
