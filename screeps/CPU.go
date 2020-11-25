package screeps

import "syscall/js"

type CPU struct {
	Limit     int
	TickLimit int
	Bucket    int
	//@todo @missing ShardLimits object<string,number>
	Unlocked     bool
	UnlockedTime int
}

func (cpu *CPU) LoadCPU(obj js.Value) {
	cpu.Limit = MarshalInt(obj.Get("limit"), 0)
	cpu.TickLimit = MarshalInt(obj.Get("tickLimit"), 0)
	cpu.Bucket = MarshalInt(obj.Get("bucket"), 0)
	cpu.Unlocked = MarshalBool(obj.Get("unlocked"), false)
	cpu.UnlockedTime = MarshalInt(obj.Get("unlockedTime"), 0)
}
