package memory

import (
	"encoding/json"
	"errors"

	"github.com/gopherjs/gopherjs/js"
)

type Memory struct {
	Rooms  map[string]*RoomMemory
	Creeps map[string]*CreepMemory
	Spawns map[string]*StructureMemory
}

type RoomMemory struct {
	Creeps []string
	Tasks  map[string]*Task
}

type Task struct {
	Name string
}

type CreepMemory struct {
	Task  string
	TaskQ []string
}

type StructureMemory struct {
}

func Get() (*Memory, error) {
	m := &Memory{
		Rooms:  map[string]*RoomMemory{},
		Creeps: map[string]*CreepMemory{},
		Spawns: map[string]*StructureMemory{},
	}
	err := m.Load()
	return m, err
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
