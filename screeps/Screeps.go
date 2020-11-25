package screeps

import "syscall/js"

func GetGame() Game {
	game := Game{}
	game.LoadGame(js.Global().Get("Game"))
	return game
}

func MarshalInt(val js.Value, def int) int {
	if val.Type() == js.TypeNumber {
		return val.Int()
	}

	return def
}

func MarshalString(val js.Value, def string) string {
	if val.Type() == js.TypeString {
		return val.String()
	}

	return def
}

func MarshalBool(val js.Value, def bool) bool {
	if val.Type() == js.TypeBoolean {
		return val.Bool()
	}

	return def
}

func MarshalEffects(obj js.Value) []Effect {
	if obj.Type() != js.TypeObject {
		return make([]Effect, 0)
	}

	effects := make([]Effect, obj.Length())
	for i := 0; i < obj.Length(); i++ {
		effect := Effect{}
		effect.LoadEffect(obj.Index(i))
		effects[i] = effect
	}

	return effects
}

func HasKeyOfType(obj js.Value, key string, jsType js.Type) bool {
	return obj.Get(key).Type() == jsType
}

func GetObjectKeys(obj js.Value) []string {
	var keys []string
	jsKeys := js.Global().Get("Object").Call("keys", obj)
	for i := 0; i < jsKeys.Length(); i++ {
		keys = append(keys, jsKeys.Index(i).String())
	}
	return keys
}
