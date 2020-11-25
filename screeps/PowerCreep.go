package screeps

import "syscall/js"

type PowerCreep struct {
	RoomObject

	ClassName  string
	DeleteTime int
	Hits       int
	HitsMax    int
	ID         string
	Level      int
	//Memory unknown @todo
	My    bool
	Name  string
	Owner Owner
	//Store store @todo
	Powers            []Power
	Saying            string
	Shard             string
	SpawnCooldownTime int
	TicksToLive       int
}

func (powerCreep *PowerCreep) LoadPowerCreep(obj js.Value) {
	powerCreep.LoadRoomObject(obj)

	powerCreep.ClassName = MarshalString(obj.Get("className"), "")
	powerCreep.DeleteTime = MarshalInt(obj.Get("deleteTime"), 0)
	powerCreep.Hits = MarshalInt(obj.Get("hits"), 0)
	powerCreep.HitsMax = MarshalInt(obj.Get("hitsMax"), 0)
	powerCreep.ID = MarshalString(obj.Get("id"), "")
	powerCreep.Level = MarshalInt(obj.Get("level"), 0)
	//powerCreep.Memory = unknown
	powerCreep.My = MarshalBool(obj.Get("my"), false)
	powerCreep.Name = MarshalString(obj.Get("name"), "")
	if HasKeyOfType(obj, "owner", js.TypeObject) {
		powerCreep.Owner.LoadOwner(obj.Get("owner"))
	}
	if HasKeyOfType(obj, "powers", js.TypeObject) {
		powerCreep.LoadPowers(obj.Get("powers"))
	}
	powerCreep.Saying = MarshalString(obj.Get("saying"), "")
	powerCreep.Shard = MarshalString(obj.Get("spawning"), "")
	powerCreep.SpawnCooldownTime = MarshalInt(obj.Get("spawnCooldownTime"), 0)
	powerCreep.TicksToLive = MarshalInt(obj.Get("ticksToLive"), 0)
}

func (powerCreep *PowerCreep) LoadPowers(obj js.Value) {
	powerCreep.Powers = make([]Power, obj.Length())
	for i := 0; i < obj.Length(); i++ {
		power := Power{}
		power.LoadPower(obj.Index(i))
		powerCreep.Powers[i] = power
	}
}
