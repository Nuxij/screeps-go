package screeps

import "syscall/js"

type Mineral struct {
	RoomObject

	Density             int
	MineralAmount       int
	MineralType         string // @todo could be constant
	ID                  string
	TicksToRegeneration int
}

func (mineral *Mineral) LoadMineral(obj js.Value) {
	mineral.LoadRoomObject(obj)

	mineral.Density = MarshalInt(obj.Get("density"), 0)
	mineral.MineralAmount = MarshalInt(obj.Get("mineralAmount"), 0)
	mineral.MineralType = MarshalString(obj.Get("mineralType"), "")
	mineral.ID = MarshalString(obj.Get("id"), "")
	mineral.TicksToRegeneration = MarshalInt(obj.Get("ticksToRegeneration"), 0)
}
