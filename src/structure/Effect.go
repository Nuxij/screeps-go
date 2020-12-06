package structure

import "github.com/gopherjs/gopherjs/js"

type Effect struct {
	*js.Object
	Effect         int `js:"effect"`
	Level          int `js:"level"`
	TicksRemaining int `js:"ticksRemaining"`
}
