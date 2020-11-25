package screeps

import "syscall/js"

type Room struct {
	Controller              StructureController
	EnergyAvailable         int
	EnergyCapacityAvailable int
	//Memory @todo
	Name     string
	Storage  StructureStorage
	Terminal StructureTerminal
	Visual   RoomVisual
}

func (room *Room) LoadRoom(obj js.Value) {
	if HasKeyOfType(obj, "controller", js.TypeObject) {
		room.Controller.LoadStructureController(obj.Get("controller"))
	}
	room.EnergyAvailable = MarshalInt(obj.Get("energyAvailable"), 0)
	room.EnergyCapacityAvailable = MarshalInt(obj.Get("energyCapacityAvailable"), 0)
	room.Name = MarshalString(obj.Get("name"), "")
	if HasKeyOfType(obj, "storage", js.TypeObject) {
		room.Storage.LoadStructureStorage(obj.Get("storage"))
	}
	if HasKeyOfType(obj, "terminal", js.TypeObject) {
		room.Terminal.LoadStructureTerminal(obj.Get("terminal"))
	}
	if HasKeyOfType(obj, "visual", js.TypeObject) {
		room.Visual.LoadRoomVisual(obj.Get("visual"))
	}
}
