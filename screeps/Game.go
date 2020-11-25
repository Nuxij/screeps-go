package screeps

import "syscall/js"

type Game struct {
	ConstructionSites map[string]ConstructionSite
	CPU               CPU
	Creeps            map[string]Creep
	Flags             map[string]Flag
	GCL               GCL
	GPL               GPL
	//Map
	//Market
	PowerCreeps map[string]PowerCreep
	//Resources
	Rooms map[string]Room
	//Shared
	Spawns     map[string]StructureSpawn
	Structures map[string]Structure
	Time       int
}

func (game *Game) LoadGame(obj js.Value) {
	if HasKeyOfType(obj, "constructionSites", js.TypeObject) {
		game.LoadConstructionSites(obj.Get("constructionSites"))
	}

	if HasKeyOfType(obj, "cpu", js.TypeObject) {
		game.CPU.LoadCPU(obj.Get("cpu"))
	}

	if HasKeyOfType(obj, "creeps", js.TypeObject) {
		game.LoadCreeps(obj.Get("creeps"))
	}

	if HasKeyOfType(obj, "flags", js.TypeObject) {
		game.LoadFlags(obj.Get("flags"))
	}

	if HasKeyOfType(obj, "gcl", js.TypeObject) {
		game.GCL.LoadGCL(obj.Get("gcl"))
	}
	if HasKeyOfType(obj, "gpl", js.TypeObject) {
		game.GPL.LoadGPL(obj.Get("gpl"))
	}

	if HasKeyOfType(obj, "powerCreeps", js.TypeObject) {
		game.LoadPowerCreeps(obj.Get("powerCreeps"))
	}

	if HasKeyOfType(obj, "rooms", js.TypeObject) {
		game.LoadRooms(obj.Get("rooms"))
	}

	if HasKeyOfType(obj, "spawns", js.TypeObject) {
		game.LoadSpawns(obj.Get("spawns"))
	}

	if HasKeyOfType(obj, "structures", js.TypeObject) {
		game.LoadStructures(obj.Get("structures"))
	}

	game.Time = MarshalInt(obj.Get("time"), 0)
}

func (game *Game) LoadConstructionSites(sites js.Value) {
	game.ConstructionSites = make(map[string]ConstructionSite)
	for _, key := range GetObjectKeys(sites) {
		site := ConstructionSite{}
		site.LoadConstructionSite(sites.Get(key))
		game.ConstructionSites[key] = site
	}
}

func (game *Game) LoadCreeps(creeps js.Value) {
	game.Creeps = make(map[string]Creep)
	for _, key := range GetObjectKeys(creeps) {
		creep := Creep{}
		creep.LoadCreep(creeps.Get(key))
		game.Creeps[key] = creep
	}
}

func (game *Game) LoadFlags(flags js.Value) {
	game.Flags = make(map[string]Flag)
	for _, key := range GetObjectKeys(flags) {
		flag := Flag{}
		flag.LoadFlag(flags.Get(key))
		game.Flags[key] = flag
	}
}

func (game *Game) LoadPowerCreeps(powerCreeps js.Value) {
	game.PowerCreeps = make(map[string]PowerCreep)
	for _, key := range GetObjectKeys(powerCreeps) {
		powerCreep := PowerCreep{}
		powerCreep.LoadPowerCreep(powerCreeps.Get(key))
		game.PowerCreeps[key] = powerCreep
	}
}

func (game *Game) LoadRooms(rooms js.Value) {
	game.Rooms = make(map[string]Room)
	for _, key := range GetObjectKeys(rooms) {
		room := Room{}
		room.LoadRoom(rooms.Get(key))
		game.Rooms[key] = room
	}
}

func (game *Game) LoadSpawns(spawns js.Value) {
	game.Spawns = make(map[string]StructureSpawn)
	for _, key := range GetObjectKeys(spawns) {
		spawn := StructureSpawn{}
		spawn.LoadStructureSpawn(spawns.Get(key))
		game.Spawns[key] = spawn
	}
}

func (game *Game) LoadStructures(structures js.Value) {
	game.Structures = make(map[string]Structure)
	for _, key := range GetObjectKeys(structures) {
		structure := Structure{}
		structure.LoadStructure(structures.Get(key))
		game.Structures[key] = structure
	}
}
