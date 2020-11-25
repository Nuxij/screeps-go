package screeps

import "syscall/js"

type GPL struct {
	Level         int
	Progress      int
	ProgressTotal int
}

func (gpl *GPL) LoadGPL(obj js.Value) {
	gpl.Level = MarshalInt(obj.Get("level"), 0)
	gpl.Progress = MarshalInt(obj.Get("progress"), 0)
	gpl.ProgressTotal = MarshalInt(obj.Get("progressTotal"), 0)
}
