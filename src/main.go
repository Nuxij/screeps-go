package main

import (
	"game"
	"logic"
	"memory"
	"pcontext"

	"github.com/gopherjs/gopherjs/js"
)

func main() {
	println("You!")
	rawMem, err := memory.Get()
	if err != nil {
		println("Fucked it memory")
		return
	}
	rawMem.Load()
	if rawMem.Options.DisableTick {
		return
	}

	pc := &pcontext.GamePointers{
		Game:   game.Form(),
		Memory: rawMem,
	}
	js.Global.Set("go", map[string]interface{}{
		"loop": func() {
			looper(pc)
		},
	})
}

func looper(pc *pcontext.GamePointers) {

	logic.RecordPreExistingCreeps(pc)

	for _, k := range logic.CollectRoomsBelowLevel(5) {
		println("Bringing Room to GCL 5: " + k.Name)
		logic.IncreaseLevel(pc, k)
	}

	PrintStats()
	pc.Memory.Save()
}

func PrintStats() {
	game := js.Global.Get("Game")
	stats := game.Get("cpu").Call("getHeapStatistics")
	ths := stats.Get("total_heap_size").Int()
	hsl := stats.Get("heap_size_limit").Int()

	println("Used ", ths, "/", hsl, "[", (float64(ths)/float64(hsl))*100, "] of the Heap\n")
}

// func ControllersNeedUpgrade() {
// game := Game.Get(js.Global().Get("Game"))
// creeps := screeps.GetObjectKeys(js.Global().Get("Game").Get("creeps"))
// fmt.Println(len(creeps))
// jimmy := game.Creeps("Jimmy")
// for roomName := range game.Rooms {
// 	controller := game.Rooms[roomName].Controller
// 	spawner := game.Spawns["Spawn1"]
// 	if controller.Progress < controller.ProgressTotal {
// 		fmt.Printf("Spawn: %+v\n", spawner)
// 		if jimmy == nil {
// 			success := spawner.SpawnCreep([]string{"move", "carry", "work"}, "Jimmy")
// 			if success == 0 {
// 				jimmy = game.Creeps["Jimmy"]
// 				println(jimmy.Name)
// 			}
// 		} else {
// 			fmt.Println("Jimmy exists")
// 			// jimmy.moveTo(spawner)
// 			//jimmy.moveTo(somePlace)
// 		}
// 	}
// }
// }

// func GetObjectKeys(obj *js.Object) []string {
// 	var keys []string
// 	jsKeys := js.Global.Get("Object").Call("keys", *obj)
// 	for i := 0; i < jsKeys.Length(); i++ {
// 		keys = append(keys, jsKeys.Index(i).String())
// 	}
// 	return keys
// }

// //export loop
// func loop() {
// 	creeps :=
// 		GetObjectKeys(js.Global().Get("Game").Get("creeps"))
// 	fmt.Println(len(creeps))
// 	// for i := 0; ; i++ { {
// 	// 	fmt.Println(n)
// 	// }
// 	//fmt.Println(js.Global().Get("Game").InstanceOf(js.New("string", "Game")))
// 	// Game := screeps.GetGame()
// 	/*
// 		fmt.Printf("CPU: %+v\n", Game.CPU)
// 		fmt.Printf("GCL: %+v\n", Game.GCL)
// 		fmt.Printf("GPL: %+v\n", Game.GPL)
// 		fmt.Printf("Creeps['dave']: %+v\n", Game.Creeps["dave"]) // will crash can tiny go is whack
// 	*/
// 	// ControllersNeedUpgrade()
// 	// println(Game.Creeps["dave"].Name)
// }
