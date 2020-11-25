package structure

import (
	"github.com/gopherjs/gopherjs/js"
)

type Structure struct {
	*js.Object

	Hits          int    `js:"hits"`
	HitsMax       int    `js:"hitsMax"`
	ID            string `js:"id"`
	StructureType string `js:"structureType"` // @todo could be constant
}
