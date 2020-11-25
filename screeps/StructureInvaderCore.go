package screeps

import "syscall/js"

type StructureInvaderCore struct {
	OwnedStructure

	Level         int
	TicksToDeploy int
}

func (structure *StructureInvaderCore) LoadStructureInvaderCore(obj js.Value) {
	structure.LoadOwnedStructure(obj)

	structure.Level = MarshalInt(obj.Get("level"), 0)
	structure.TicksToDeploy = MarshalInt(obj.Get("ticksToDeploy"), 0)
}
