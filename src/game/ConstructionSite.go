package game

import "github.com/gopherjs/gopherjs/js"

type ConstructionSite struct {
	*js.Object

	ID string `js:id`
	My bool   `js:my`
	// Owner         *Owner `js:owner`
	Progress      int `js:progress`
	ProgressTotal int `js:progressTotal`
	// StructureType string `js:structureType`
}
