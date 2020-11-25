package screeps

import "syscall/js"

type StructurePortal struct {
	Structure

	Destination      RoomPosition
	DestinationShard ShardPosition
	TicksToDecay     int
}

func (structure *StructurePortal) LoadStructurePortal(obj js.Value) {
	structure.LoadStructure(obj)

	if HasKeyOfType(obj, "destination", js.TypeObject) {
		structure.Destination.LoadRoomPosition(obj.Get("destination"))
		structure.DestinationShard.LoadShardPosition(obj.Get("destination"))
	}
	structure.TicksToDecay = MarshalInt(obj.Get("ticksToDecay"), 0)
}
