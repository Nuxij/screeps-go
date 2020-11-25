package screeps

import "syscall/js"

type Sign struct {
	Username string
	Text     string
	Time     int
	//DateTime @todo should be a Date object?
}

func (sign *Sign) LoadSign(obj js.Value) {
	sign.Username = MarshalString(obj.Get("username"), "")
	sign.Text = MarshalString(obj.Get("text"), "")
	sign.Time = MarshalInt(obj.Get("time"), 0)
}
