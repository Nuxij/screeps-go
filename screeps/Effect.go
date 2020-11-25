package screeps

import "syscall/js"

type Effect struct {
	Effect         int
	Level          int
	TicksRemaining int
}

func (effect *Effect) LoadEffect(obj js.Value) {
	effect.Effect = MarshalInt(obj.Get("effect"), 0)
	effect.Level = MarshalInt(obj.Get("level"), 0)
	effect.TicksRemaining = MarshalInt(obj.Get("ticksRemaining"), 0)
}
