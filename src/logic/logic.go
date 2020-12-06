package logic

import (
	"enum"
	"game"
	"math"
	"room"
	"structure"

	"github.com/gopherjs/gopherjs/js"
)

func Stringify(o *js.Object) string {
	return js.Global.Get("JSON").Call("stringify", o).String()
}

func findInArray(s []string, element string) bool {
	for i := range s {
		println(s[i])
		if s[i] == element {
			return true
		}
	}
	return false
}

func CollectRoomsBelowLevel(level int) []*room.Room {
	g := game.Form()
	println("Finding rooms", level)
	var rooms []*room.Room
	for _, k := range g.Rooms {
		if k.Controller.Level < level {
			rooms = append(rooms, k)
		}
	}
	return rooms
}

// func SpawnCreep(r *room.Room) bool {
// 	spawn := FindSpawnInRoom(r)
// 	if spawn == nil {
// 		println("Spawn could not be found in room " + r.Name)
// 		return false
// 	}

// 	name := utils.RandoName()
// 	spawn.Call("spawnCreep", []string{consts.MOVE, consts.WORK, consts.CARRY}, name)

// 	mem, err := memory.Get()
// 	if err != nil {
// 		return false
// 	}

// 	rmem, ok := mem.Rooms[r.Name]
// 	if !ok {
// 		rmem = &memory.RoomMemory{}
// 		mem.Rooms[r.Name] = rmem
// 	}

// 	ok = findInArray(rmem.Creeps, name)
// 	if !ok {
// 		rmem.Creeps = append(rmem.Creeps, name)
// 	}

// 	err = mem.Save()
// 	if err != nil {
// 		println(err.Error())
// 		return false
// 	}

// 	return true
// }

// func SpawnCreepsForRoom(r *room.Room) {
// 	creeps := GetCreepsInRoom(r.Name)
// 	println(len(creeps))
// 	if len(creeps) < 2 {
// 		SpawnCreep(r)
// 	}
// }

func FindEnergyDestinations(r *room.Room) []*structure.Structure {
	d := []*structure.Structure{}
	ss := r.Call("find", enum.FIND_MY_STRUCTURES, func(s *structure.Structure) bool {
		st := s.StructureType
		return st == enum.STRUCTURE_CONTAINER || st == enum.STRUCTURE_SPAWN
	})
	for i := 0; i < ss.Length(); i++ {
		d = append(d, &structure.Structure{RoomObject: &room.RoomObject{Object: ss.Index(i)}})
	}
	return d
}

func StoreEnergy(r *room.Room, max int) {
	maxContainersPerRoom := 5 // hard limit in the game

	// Find source
	// Find appropriate destination
	// If full, find container
	// no containers, create one
	// can't do anything, get a new task
	// source := FindEnergySources(r)
	dests := FindEnergyDestinations(r)
	containers := 0
	for _, d := range dests {
		// Store task "StoreEnergy, s, d"
		// Only on destinations that have enough room (> 100?)
		if d.StructureType == enum.STRUCTURE_CONTAINER {
			containers++
		}
	}

	println("  Ensuring containers")
	if containers < maxContainersPerRoom {
		optimalContainers := math.Ceil(float64(max) / 2000)
		println(optimalContainers)
		// Create a single container
		// Or ask for it's task or w.e
	}

	// SpawnCreepsForRoom(r)
}
