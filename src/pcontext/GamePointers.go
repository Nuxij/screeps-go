package pcontext

import (
	"game"
	"memory"
)

type GamePointers struct {
	Game   *game.Game
	Memory *memory.Memory
}
