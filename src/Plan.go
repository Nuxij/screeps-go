package main

import "context"

// Plan describe a set of procedures and the way to calculate them
type Plan interface {
	// Create returns the list of procedures that needs to be executed.
	Create(ctx context.Context) ([]Procedure, error)
	// Name identifies a specific plan
	Name() string
}

type RoomUpgradePlan struct {
	Room string
	TargetGCL int
}

// func FindEnergySources(r string) *game.Source {
// 	ss := r.Call("find", consts.FIND_SOURCES_ACTIVE, func(s *game.Source) bool { return s.Room.Name == r.Name })
// 	return &game.Source{RoomObject: &room.RoomObject{Object: ss}}
// }

func (rup *RoomUpgradePlan) Create(ctx context.Context) ([]Procedure, error) {
	steps := []Procedure{
		GetEnergy{
			amount: 50,
			Source: 
		}
}