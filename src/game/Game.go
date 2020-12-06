package game

import (
	"creep"
	"room"
	"structure"

	"github.com/gopherjs/gopherjs/js"
)

type Game struct {
	*js.Object
	Rooms             map[string]*room.Room        `js:"rooms"`
	ConstructionSites map[string]*ConstructionSite `js:"constructionSites"`
	Time              int                          `js:"time"`
	Creeps            map[string]*creep.Creep      `js:"creeps"`
	Spawns            map[string]*structure.Spawn  `js:"spawns"`
}

func Form() *Game {
	return &Game{Object: js.Global.Get("Game")}
}
