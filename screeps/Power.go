package screeps

import "syscall/js"

type Power struct {
	Level    int
	Cooldown int
}

func (power *Power) LoadPower(obj js.Value) {
	power.Level = MarshalInt(obj.Get("level"), 0)
	power.Cooldown = MarshalInt(obj.Get("cooldown"), 0)
}
