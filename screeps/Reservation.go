package screeps

import "syscall/js"

type Reservation struct {
	Username   string
	TicksToEnd int
}

func (reservation *Reservation) LoadReservation(obj js.Value) {
	reservation.Username = MarshalString(obj.Get("username"), "")
	reservation.TicksToEnd = MarshalInt(obj.Get("ticksToEnd"), 0)
}
