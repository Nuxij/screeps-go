package screeps

import "syscall/js"

type StructureController struct {
	OwnedStructure

	IsPowerEnabled    bool
	Level             int
	Progress          int
	ProgressTotal     int
	Reservation       Reservation
	SafeMode          int
	SafeModeAvailable int
	SafeModeCooldown  int
	Sign              Sign
	TicksToDowngrade  int
	UpgradeBlocked    int
}

func (structure *StructureController) LoadStructureController(obj js.Value) {
	structure.LoadOwnedStructure(obj)

	structure.IsPowerEnabled = MarshalBool(obj.Get("isPowerEnabled"), false)
	structure.Level = MarshalInt(obj.Get("level"), 0)
	structure.Progress = MarshalInt(obj.Get("progress"), 0)
	structure.ProgressTotal = MarshalInt(obj.Get("progressTotal"), 0)
	if HasKeyOfType(obj, "reservation", js.TypeObject) {
		structure.Reservation.LoadReservation(obj.Get("reservation"))
	}
	structure.SafeMode = MarshalInt(obj.Get("safeMode"), 0)
	structure.SafeModeAvailable = MarshalInt(obj.Get("safeModeAvailable"), 0)
	structure.SafeModeCooldown = MarshalInt(obj.Get("safeModeCooldown"), 0)
	if HasKeyOfType(obj, "sign", js.TypeObject) {
		structure.Sign.LoadSign(obj.Get("sign"))
	}
	structure.TicksToDowngrade = MarshalInt(obj.Get("ticksToDowngrade"), 0)
	structure.UpgradeBlocked = MarshalInt(obj.Get("upgradeBlocked"), 0)
}
