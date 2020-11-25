package room

import (
	"github.com/gopherjs/gopherjs/js"
)

type RoomObject struct {
	*js.Object
	// Effects []*game.Effect `js:"effects"`
	Pos  *RoomPosition `js:"pos"`
	Room *Room         `js:"room"`
}
