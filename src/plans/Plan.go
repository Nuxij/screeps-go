package plans

import (
	"pcontext"
	"procedures"
)

// Plan describe a set of procedures and the way to calculate them
type Plan interface {
	// Create returns the list of procedures that needs to be executed.
	Create(pc *pcontext.GamePointers) ([]procedures.Procedure, error)
	// Name identifies a specific plan
	Name() string
}
