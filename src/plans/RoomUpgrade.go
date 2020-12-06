package plans

import (
	"enum"
	"memory"
	"pcontext"
	"procedures"
	"room"
	"strconv"
)

type RoomUpgradePlan struct {
	Room      string
	TargetGCL int
}

func Sources(r *room.Room) []*room.Source {
	sources := []*room.Source{}
	filter := func(s *room.Source) bool { return s.Room.Name == r.Name }
	res := r.Call("find", enum.FIND_SOURCES_ACTIVE, filter)
	for si := 0; si < res.Length(); si++ {
		sources = append(sources, &room.Source{RoomObject: &room.RoomObject{Object: res.Index(si)}})
	}
	return sources
}

func (rup RoomUpgradePlan) Name() string {
	return "room-upgrade-" + rup.Room + "-" + strconv.Itoa(rup.TargetGCL)
}

func (rup RoomUpgradePlan) Create(pc *pcontext.GamePointers) ([]procedures.Procedure, error) {
	steps := []procedures.Procedure{}
	roomForUpgrade := room.Form(rup.Room)

	mem, ok := pc.Memory.Rooms[rup.Room]
	if !ok {
		mem = &memory.RoomMemory{
			Tasks: []string{rup.Name()},
		}
		pc.Memory.Rooms[rup.Room] = mem
	}

	c := roomForUpgrade.Controller
	energyToNextLevel := c.ProgressTotal - c.Progress

	sources := Sources(roomForUpgrade)
	for _, source := range sources {
		steps = append(steps, procedures.GetEnergy{
			Amount: energyToNextLevel / len(sources),
			Source: source,
			Room:   roomForUpgrade,
		})
	}

	// steps = append(steps, procedures.PutEnergy{
	// 	Amount: energyToNextLevel,
	// 	Room:   roomForUpgrade,
	// 	Target: roomForUpgrade.Controller.RoomObject,
	// })

	// spawn := &structure.Spawn{
	// 	Structure: &structure.Structure{
	// 		RoomObject: &room.RoomObject{
	// 			Object: roomForUpgrade.Call("find", consts.FIND_MY_SPAWNS).Index(0),
	// 		},
	// 	},
	// }

	// g := game.Form()
	// if len(g.Creeps) < 2 {
	// 	steps = append(steps, procedures.SpawnCreep{Spawn: spawn})
	// }

	return steps, nil
}
