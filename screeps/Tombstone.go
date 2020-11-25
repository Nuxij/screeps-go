package screeps

import "syscall/js"

type Tombstone struct {
	RoomObject

	Creep        Creep
	PowerCreep   PowerCreep
	DeathTime    int
	ID           string
	Store        Store
	TicksToDecay int
}

func (tombstone *Tombstone) LoadTombstone(obj js.Value) {
	tombstone.LoadRoomObject(obj)

	if HasKeyOfType(obj, "creep", js.TypeObject) {
		tombstone.Creep.LoadCreep(obj.Get("creep"))
		tombstone.PowerCreep.LoadPowerCreep(obj.Get("creep"))
	}
	tombstone.DeathTime = MarshalInt(obj.Get("deathTime"), 0)
	tombstone.ID = MarshalString(obj.Get("id"), "")
	if HasKeyOfType(obj, "store", js.TypeObject) {
		tombstone.Store.LoadStore(obj.Get("store"))
	}
	tombstone.TicksToDecay = MarshalInt(obj.Get("ticksToDecay"), 0)
}
