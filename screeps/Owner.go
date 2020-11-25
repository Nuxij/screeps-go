package screeps

import "syscall/js"

type Owner struct {
	Username string
}

func (owner *Owner) LoadOwner(obj js.Value) {
	owner.Username = MarshalString(obj.Get("username"), "")
}
