package screeps

import "syscall/js"

type GCL struct {
	Level         int
	Progress      int
	ProgressTotal int
}

func (gcl *GCL) LoadGCL(obj js.Value) {
	gcl.Level = MarshalInt(obj.Get("level"), 0)
	gcl.Progress = MarshalInt(obj.Get("progress"), 0)
	gcl.ProgressTotal = MarshalInt(obj.Get("progressTotal"), 0)
}
