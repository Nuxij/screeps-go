package screeps

import "syscall/js"

type StructurePowerBank struct {
	Structure

	Power        int
	TicksToDecay int
}

func (structure *StructurePowerBank) LoadStructurePowerBank(obj js.Value) {
	structure.LoadStructure(obj)

	structure.Power = MarshalInt(obj.Get("power"), 0)
	structure.TicksToDecay = MarshalInt(obj.Get("ticksToDecay"), 0)
}
