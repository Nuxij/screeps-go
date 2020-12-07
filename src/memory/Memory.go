package memory

import (
	"encoding/json"
	"errors"

	"github.com/gopherjs/gopherjs/js"
)

type StringKey string

var CTXMemory = StringKey("memory")

type Memory struct {
	JUNK    string
	Options *OptionsMemory
	Rooms   map[string]*RoomMemory
	Creeps  map[string]*CreepMemory
}

type OptionsMemory struct {
	DisableTick bool
}

type RoomMemory struct {
	Creeps []string
	Tasks  []string
}

type Task struct {
	Name string
}

type CreepMemory struct {
	State string
	Task  string
	TaskQ []string
}

func Get() (*Memory, error) {
	m := &Memory{
		JUNK: "just junk",
		Options: &OptionsMemory{
			DisableTick: false,
		},
		Rooms:  map[string]*RoomMemory{},
		Creeps: map[string]*CreepMemory{},
	}
	return m, nil
}

func (m *Memory) Load() error {
	data := js.Global.Get("RawMemory").Call("get")
	if data == nil {
		return errors.New("RawMemory not found!")
	}
	err := json.Unmarshal([]byte(data.String()), m)
	if err != nil {
		return err
	}
	return nil
}

func (m *Memory) Save() error {
	jsonBytes, err := json.Marshal(m)
	if err != nil {
		return err
	}
	js.Global.Get("RawMemory").Call("set", string(jsonBytes))
	return nil
}

// func NewRoom(name string) *RoomMemory {
// 	return &RoomMemory{
// 		Object: js.Global.Get("Object").New(),
// 		Creeps: js.Global.Get("Array").New(),
// 		Name:   name,
// 	}
// }
