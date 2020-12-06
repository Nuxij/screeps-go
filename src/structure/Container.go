package structure

type Container struct {
	*Structure

	// Store        *Store
	TicksToDecay int `js:"ticksToDecay"`
}
