package room

type Source struct {
	*RoomObject
	Energy              int    `js:"energy"`
	EnergyCapacity      int    `js:"energyCapacity"`
	ID                  string `js:"id"`
	TicksToRegeneration int    `js:"ticksToRegeneration"`
}
