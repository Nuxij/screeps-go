package main

import "github.com/gopherjs/gopherjs/js"

func main() {
	println(js.Global.Get("Game").Get("spawns").Get("Spawn1"))
}

func loop() {
	println("Yo")
}
