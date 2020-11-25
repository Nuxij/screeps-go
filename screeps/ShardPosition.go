package screeps

import "syscall/js"

type ShardPosition struct {
	Shard string
	Room  string
}

func (pos *ShardPosition) LoadShardPosition(obj js.Value) {
	pos.Shard = MarshalString(obj.Get("shard"), "")
	pos.Room = MarshalString(obj.Get("room"), "")
}
