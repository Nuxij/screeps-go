package structure

import (
	"room"
)

type ConstructionSite struct {
	*room.RoomObject

	Progress      int    `js:"progress"`
	ProgressTotal int    `js:"progressTotal"`
	ID            string `js:"id"`
	StructureType string `js:"structureType"` // @todo could be constant
	// Effects       []*Effect `js:"effects"`
}
