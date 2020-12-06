package logic

import (
	"memory"
	"pcontext"
)

func RecordPreExistingCreeps(pc *pcontext.GamePointers) {
	for name, _ := range pc.Memory.Creeps {
		_, ok := pc.Game.Creeps[name]
		if !ok {
			delete(pc.Memory.Creeps, name)
		}
	}
	for _, c := range pc.Game.Creeps {
		_, ok := pc.Memory.Creeps[c.Name]
		if !ok {
			pc.Memory.Creeps[c.Name] = &memory.CreepMemory{
				Task: "none",
			}
		}
	}
}
