package utils

import (
	"errors"

	"github.com/gopherjs/gopherjs/js"
)

type FSMHandler interface {
	Transition()
}

type FSMHandlerFunc func(FSMHandlerFunc) FSMHandlerFunc

// func (def *FSMDefinition) Transition(name string) {
//   for from, to := range def.Transitions {
//     if
//   }
// }

type FSM struct {
	state string

	InitialState string
	States       map[string]map[string]FSMHandlerFunc
}

func (fsm *FSM) New(handler *FSMHandler, currentState string) {
	// fsm.handler = handler
	// state = currentState
}

func (fsm *FSM) Transition(name string) {
	// const spec = this.transitions().find(t => t.name == transition)
	// if(!spec)
	//   throw new Error(`Cannot use transition '${transition}' from state '${this.state}'.`)

	// const to = this._states.find(s => s.name == spec.to)
	// const from = this._states.find(s => s.name == spec.from)

	// this.trigger("Before", spec)
	// this.trigger("Leave", from)

	// this.state = spec.to

	// this.trigger("Enter", to)
	// this.trigger("After", spec)

}

func (fsm *FSM) Run() error {
	if js.Global.Get("Memory").Get("forceResetFSM").Bool() == true {
		fsm.state = fsm.InitialState
	}

	_, ok := fsm.States[fsm.state]
	if !ok {
		return errors.New("Cannot find state in computed states list.")
	}

	// fsm.Transition("", transitionSpec)
	return nil

	// if(typeof this.onAny == "function") {}
	// 	return this.onAny.call(this.handler, this)
	// }
}

//       if(this.state !== undefined)
//         this.state = state
//       this.state = this.machine.init

//       this._states = []
//       this._transitions = this.machine.transitions

//       this._transitions.forEach(t => {
//         if(!t.camelName)
//           t.camelName = t.name.substr(0, 1).toUpperCase() + t.name.substr(1)

//         this[t.name] = () => this.transition(t.name)
//         if(!this._states.find(s => s.name == t.to))
//           this._states.push({ name: t.to })
//         if(!this._states.find(s => s.name == t.from))
//           this._states.push({ name: t.from })
//       })

//       this._states.forEach(s => {
//         s.camelName = s.name.substr(0, 1).toUpperCase() + s.name.substr(1)
//       })
//     }

//     run() {
//       if(Memory.forceResetFSM == true)
//         this.state = this.machine.init

//       const spec = fsm.States.find(s => s.name == this.state)
//       if(!spec)
//         throw new Error(`Cannot find state '${this.state}' in computed states list.`)

//       this.trigger("", spec)

//       if(typeof this.onAny == "function")
//         return this.onAny.call(this.handler, this)
//     }

//     transition(transition) {

//     }

//     is(state) {
//       return this.state == state
//     }

//     can(state) {
//       return this.transitions().indexOf(state) !== -1
//     }

//     cannot(state) {
//       return !this.can(state)
//     }

//     transitions() {
//       return this._transitions.filter(t => t.from == this.state || t.from == "any")
//     }

//     allTransitions() {
//       return this._transitions
//     }

//     allStates() {
//       return this._states
//     }

//     trigger(eventName, spec) {
//       if(typeof this[`on${eventName}${spec.camelName}`] == "function")
//         return this[`on${eventName}${spec.camelName}`].call(this.handler, this)
//     }
// }
