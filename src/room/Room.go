package room

import (
	"github.com/gopherjs/gopherjs/js"
)

type Room struct {
	*js.Object
	Name            string      `js:"name"`
	Controller      *Controller `js:"controller"`
	EnergyAvailable int         `js:"energyAvailable"`
}

func Form(id string) *Room {
	g := js.Global.Get("Game")
	r := g.Get("rooms").Get(id)
	if r == nil {
		return nil
	}
	return &Room{Object: r}
}

// type RoomWatch struct {
// 	*js.Object
// 	Processed []string `js:processed`
// 	Busy      []string `js:busy`
// }
