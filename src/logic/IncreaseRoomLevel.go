package logic

import (
	"pcontext"
	"plans"
	"room"
)

func IncreaseLevel(pc *pcontext.GamePointers, r *room.Room) error {
	upgradePlan := plans.RoomUpgradePlan{
		Room:      r.Name,
		TargetGCL: 5,
	}

	scheduler := &Scheduler{}

	if err := scheduler.Execute(pc, upgradePlan); err != nil {
		println(err.Error())
		return err
	}

	return nil
}
