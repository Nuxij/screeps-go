package screeps

import "syscall/js"

type Source struct {
	RoomObject

	Energy              int
	EnergyCapacity      int
	ID                  string
	TicksToRegeneration int
}

func (source *Source) LoadRuin(obj js.Value) {
	source.LoadRoomObject(obj)

	source.Energy = MarshalInt(obj.Get("energy"), 0)
	source.EnergyCapacity = MarshalInt(obj.Get("energyCapacity"), 0)
	source.ID = MarshalString(obj.Get("id"), "")
	source.TicksToRegeneration = MarshalInt(obj.Get("ticksToRegeneration"), 0)
}
