package screeps

import "syscall/js"

type Nuke struct {
	RoomObject

	ID             string
	LaunchRoomName string
	TimeToLand     int
}

func (nuke *Nuke) LoadNuke(obj js.Value) {
	nuke.LoadRoomObject(obj)

	nuke.ID = MarshalString(obj.Get("id"), "")
	nuke.LaunchRoomName = MarshalString(obj.Get("launchRoomName"), "")
	nuke.TimeToLand = MarshalInt(obj.Get("timeToLand"), 0)
}
