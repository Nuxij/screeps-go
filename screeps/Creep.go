package screeps

import "syscall/js"

type Creep struct {
	RoomObject

	Name    string
	Body    []CreepBodyPart
	Fatigue int
	Hits    int
	HitsMax int
	ID      string
	//Memory unknown @todo
	My       bool
	Owner    Owner
	Saying   string
	Spawning bool
	//Store store @todo
	TicksToLive int
}

func (creep *Creep) LoadCreep(obj js.Value) {
	creep.LoadRoomObject(obj)

	if HasKeyOfType(obj, "body", js.TypeObject) {
		creep.LoadBodyParts(obj.Get("body"))
	}

	creep.Fatigue = MarshalInt(obj.Get("fatigue"), 0)
	creep.Hits = MarshalInt(obj.Get("hits"), 0)
	creep.HitsMax = MarshalInt(obj.Get("hitsMax"), 0)
	creep.ID = MarshalString(obj.Get("id"), "")
	//creep.Memory = unknown @todo
	creep.My = MarshalBool(obj.Get("my"), false)
	creep.Owner.LoadOwner(obj.Get("owner"))
	creep.Saying = MarshalString(obj.Get("saying"), "")
	creep.Spawning = MarshalBool(obj.Get("spawning"), false)
	//creep.Store = store @todo
	creep.TicksToLive = MarshalInt(obj.Get("ticksToLive"), 0)
	creep.Name = MarshalString(obj.Get("name"), "")
}

func (creep *Creep) LoadBodyParts(obj js.Value) {
	creep.Body = make([]CreepBodyPart, obj.Length())
	for i := 0; i < obj.Length(); i++ {
		part := CreepBodyPart{}
		part.LoadCreepBodyPart(obj.Index(i))
		creep.Body[i] = part
	}
}
