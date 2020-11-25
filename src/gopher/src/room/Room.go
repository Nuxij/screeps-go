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

// func Form(id string) Room {
// 	room := Room{}
// 	room.Game = js.Global.Get("Game")
// 	room.Obj = room.Game.Get("rooms").Get(id)
// 	room.Name = room.Obj.Get("name").String()
// 	return room
// }
