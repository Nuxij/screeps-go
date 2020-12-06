package game

import "creep"

func FindCreepsInRoom(name string) []*creep.Creep {
	g := Form()
	var creeps []*creep.Creep
	for _, k := range g.Creeps {
		if k.Room.Name == name {
			creeps = append(creeps, k)
		}
	}
	return creeps
}
