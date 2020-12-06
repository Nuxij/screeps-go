package structure

import (
	"room"
)

type Structure struct {
	*room.RoomObject

	Hits          int    `js:"hits"`
	HitsMax       int    `js:"hitsMax"`
	ID            string `js:"id"`
	StructureType string `js:"structureType"` // @todo could be constant
	// Effects       []*Effect `js:"effects"`
}
