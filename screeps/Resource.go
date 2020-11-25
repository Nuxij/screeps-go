package screeps

import "syscall/js"

type Resource struct {
	RoomObject

	Amount       int
	ID           string
	ResourceType string // @todo could be a constance
}

func (resource *Resource) LoadResource(obj js.Value) {
	resource.LoadRoomObject(obj)

	resource.Amount = MarshalInt(obj.Get("className"), 0)
	resource.ID = MarshalString(obj.Get("id"), "")
	resource.ResourceType = MarshalString(obj.Get("resourceType"), "")
}
