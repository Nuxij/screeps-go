package game

import (
	"github.com/gopherjs/gopherjs/js"
)

type CreepBodyPart struct {
	*js.Object
	Boost string `js:"boost"`
	Type  string `js:"type"`
	Hits  int    `js:"hits"`
}
