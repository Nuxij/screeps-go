package game

import (
	"enum"
	"room"
	"structure"

	"github.com/gopherjs/gopherjs/js"
)

func FindSpawnsInRoom(name string) []*structure.Spawn {
	g := Form()
	// ss := r.Call("find", consts.FIND_MY_SPAWNS, func(s *structure.Spawn) bool { return s.Room.Name == r.Name })
	// spawn, ok := g.Spawns[ss.Index(0).Get("name").String()]
	var spawns []*structure.Spawn
	for _, s := range g.Spawns {
		if s.Room.Name == name {
			spawns = append(spawns, s)
		}
	}
	return spawns
}

func FindContainerAtPos(p *room.RoomPosition) *structure.Container {
	g := Form()
	r := g.Rooms[p.RoomName]
	obj := r.Call("find", enum.FIND_MY_STRUCTURES, func(s *structure.Structure) bool { return s.StructureType == "container" }) //s.Pos.X == p.X && s.Pos.Y == p.Y })
	println(js.Global.Get("JSON").Call("stringify", obj).String())
	if obj == nil || obj.Bool() == false || obj.Int() == 0 {
		return nil
	}
	return &structure.Container{
		Structure: &structure.Structure{
			RoomObject: &room.RoomObject{
				Object: obj,
			},
		},
	}
}
