package procedures

import (
	"pcontext"

	"github.com/gopherjs/gopherjs/js"
)

// Procedure describe every single step to be executed. It is the smallest unit
// of work in a plan.
type Procedure interface {
	// Name identifies a specific procedure.
	Name() string
	// Do execute the business logic for a specific procedure.
	Do(pc *pcontext.GamePointers) ([]Procedure, error)
}

func Stringify(o *js.Object) string {
	return js.Global.Get("JSON").Call("stringify", o).String()
}
