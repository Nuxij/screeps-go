package screeps

import "syscall/js"

type ConstructionSite struct {
	RoomObject

	ID            string
	My            bool
	Owner         Owner
	Progress      int
	ProgressTotal int
	StructureType string // @todo could be a proper constant
}

func (site *ConstructionSite) LoadConstructionSite(obj js.Value) {
	site.LoadRoomObject(obj)

	site.ID = obj.Get("id").String()
	site.My = obj.Get("my").Bool()
	if HasKeyOfType(obj, "owner", js.TypeObject) {
		site.Owner.LoadOwner(obj.Get("owner"))
	}
	site.Progress = MarshalInt(obj.Get("progress"), 0)
	site.ProgressTotal = MarshalInt(obj.Get("progressTotal"), 0)
	site.StructureType = MarshalString(obj.Get("structureType"), "")
}
