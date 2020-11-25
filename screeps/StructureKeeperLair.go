package screeps

import "syscall/js"

type StructureKeeperLair struct {
	OwnedStructure

	TicksToSpawn int
}

func (structure *StructureKeeperLair) LoadStructureKeeperLair(obj js.Value) {
	structure.LoadOwnedStructure(obj)

	structure.TicksToSpawn = MarshalInt(obj.Get("ticksToSpawn"), 0)
}
