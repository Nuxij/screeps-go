package procedures

import (
	"enum"
	"errors"
	"game"
	"memory"
	"pcontext"
	"room"
	"structure"
	"utils"
)

type SpawnCreep struct {
	CreepName     string
	Spawn         *structure.Spawn
	InitialMemory *memory.CreepMemory
}

func (s SpawnCreep) Name() string {
	return "spawn-creep-" + s.CreepName
}

func (s SpawnCreep) Do(pc *pcontext.GamePointers) ([]Procedure, error) {
	err := s.Spawn.Call("spawnCreep", []string{enum.MOVE, enum.WORK, enum.CARRY}, s.CreepName).Int()
	namedErr := enum.Errors.Name(err)
	switch namedErr {
	case "ERR_OK":
		println("OK!")
		pc.Memory.Creeps[s.CreepName] = s.InitialMemory
	default:
		println("Failed Creating Creep! " + namedErr)
	}
	return []Procedure{}, errors.New(namedErr)
}

func SpawnNamedCreep(r *room.Room, im *memory.CreepMemory) (SpawnCreep, error) {
	if im == nil {
		im = &memory.CreepMemory{}
	}
	spawnInRoom := game.FindSpawnsInRoom(r.Name)
	if len(spawnInRoom) == 0 {
		return SpawnCreep{}, errors.New("Found no spawns in " + r.Name)
	}
	return SpawnCreep{
		CreepName:     utils.RandoName(),
		Spawn:         spawnInRoom[0],
		InitialMemory: im,
	}, nil
}
