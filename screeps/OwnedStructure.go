package screeps

import "syscall/js"

type OwnedStructure struct {
	Structure

	My    bool
	Owner Owner
}

func (structure *OwnedStructure) LoadOwnedStructure(obj js.Value) {
	structure.LoadStructure(obj)

	structure.My = MarshalBool(obj.Get("my"), false)

	if HasKeyOfType(obj, "owner", js.TypeObject) {
		structure.Owner.LoadOwner(obj.Get("owner"))
	}
}
