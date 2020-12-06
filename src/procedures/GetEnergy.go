package procedures

import (
	"creep"
	"enum"
	"game"
	"memory"
	"pcontext"
	"room"
	"strconv"
	"structure"

	"github.com/pkg/errors"
)

type GetEnergy struct {
	Amount int
	Room   *room.Room
	Source *room.Source
}

func (s GetEnergy) Name() string {
	return "get-energy-" + s.Source.ID + "-" + strconv.Itoa(s.Amount)
}

func (s GetEnergy) Do(pc *pcontext.GamePointers) ([]Procedure, error) {
	steps := []Procedure{}

	creep, creepMemory := s.GetCreepForTask(pc, s.Room.Name)

	if creep == nil && creepMemory == nil {
		step, err := SpawnNamedCreep(s.Room, &memory.CreepMemory{
			Task: s.Name(),
		})
		if err != nil {
			return steps, errors.New("Can't get energy with no creeps")
		}
		steps = append(steps, step)
		return steps, nil
	}

	// OK! NOW ITS TIME FOR AN FSM!
	store := creep.Get("store")
	if store.Get("energy").Int() != store.Call("getCapacity").Int() {
		err := s.Harvest(creep)
		if err != nil {
			return steps, err
		}
	} else {
		containers := creep.Pos.Call("findInRange", enum.FIND_MY_STRUCTURES, 0)
		if containers.Get("length").Int() == 0 {
			sites := creep.Pos.Call("findInRange", enum.FIND_CONSTRUCTION_SITES, 0)
			if sites.Get("length").Int() == 0 {
				println(enum.Errors.Name(creep.Pos.Call("createConstructionSite", "container").Int()))
			} else {
				site := &structure.ConstructionSite{
					RoomObject: &room.RoomObject{
						Object: sites.Get("0"),
					},
				}
				println(site.Progress)
				creep.Call("build", sites.Get("0"))
			}
		} else {
			creep.Call("drop", "energy")
		}
		// Make a container if there isn't one
		// Drop the shit on the floor
	}

	return steps, nil
}

func (s GetEnergy) GetCreepForTask(pc *pcontext.GamePointers, roomName string) (*creep.Creep, *memory.CreepMemory) {
	var ourGuy *creep.Creep
	var ourMemory *memory.CreepMemory

	creeps := game.FindCreepsInRoom(roomName)

	for _, c := range creeps {
		cmem, ok := pc.Memory.Creeps[c.Name]
		if !ok {
			return nil, nil
		}

		if cmem.Task == "" || cmem.Task == "none" {
			cmem.Task = s.Name()
		}

		if cmem.Task == s.Name() {
			ourGuy = c
			ourMemory = cmem
			break
		}
	}
	return ourGuy, ourMemory
}

func (s GetEnergy) Harvest(creep *creep.Creep) error {
	tryHarvest := enum.Errors.Name(creep.Harvest(s.Source.RoomObject))
	switch tryHarvest {
	case "ERR_OK":
		println("Harvested Successfully!")
		return nil
	case "ERR_NO_BODYPART":
		return errors.New("Needs more CARRY")
	case "ERR_NOT_IN_RANGE":
		creep.MoveTo(s.Source.RoomObject.Pos)
		return nil
	case "ERR_INVALID_TARGET":
		return errors.New("Get-Energy given a bad source " + s.Source.ID)
	default:
		return errors.New("[Harvest:" + s.Source.ID + "] FAILED! " + tryHarvest)
	}
}
