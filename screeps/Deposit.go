package screeps

import "syscall/js"

type Deposit struct {
	RoomObject

	Cooldown     int
	DepositType  string
	ID           string
	LastCooldown int
	TicksToDecay int
}

func (deposit *Deposit) LoadDeposit(obj js.Value) {
	deposit.LoadRoomObject(obj)

	deposit.Cooldown = MarshalInt(obj.Get("cooldown"), 0)
	deposit.DepositType = MarshalString(obj.Get("depositType"), "")
	deposit.ID = MarshalString(obj.Get("id"), "")
	deposit.LastCooldown = MarshalInt(obj.Get("lastCooldown"), 0)
	deposit.TicksToDecay = MarshalInt(obj.Get("ticksToDecay"), 0)
}
