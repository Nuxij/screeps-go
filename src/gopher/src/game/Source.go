package game

import (
	"room"
)

type Source struct {
	*room.RoomObject
	Energy              int    `js:"energy"`
	EnergyCapacity      int    `js:"energyCapacity"`
	ID                  string `js:"id"`
	TicksToRegeneration int    `js:"ticksToRegeneration"`
}
