package ast

import "fmt"

const maxStates = 255

type Program struct {
	name        string
	states      map[string]State
	initState   string
	acceptState string
	currState   string
}

// GetStatesCopy clones the states of the program
func (program *Program) GetStatesCopy() map[string]State {
	statesClone := make(map[string]State, len(program.states))
	for k, v := range program.states {
		statesClone[k] = v
	}

	return statesClone
}

func (program *Program) getStateByName(name string) State {

	_, ok := program.states[name]
	if ok {
		return program.states[name]
	} else {
		program.states[name] = NewState(name)
	}

	return program.states[name]
}

//AddRule appends a rule to the program
func (program *Program) AddRule(frState, toState, frChar, toChar string, dir string) {
	instruction := NewTapeInstruction(toChar, dir)
	action := NewAction(toState, instruction)
	state := program.getStateByName(frState)
	state.Mappings[frChar] = action
}

// AddName updates name information about the program
func (program *Program) AddName(name string) {
	program.name = name
}

// AddInitState updates initial state of the program
func (program *Program) AddInitState(init string) {
	program.initState = init
	program.currState = init
}

// AddAcceptState updates the accpeting state of the program
func (program *Program) AddAcceptState(accept string) {
	program.acceptState = accept
}

// GetAcceptState returns the accept state
func (program *Program) GetAcceptState() string {
	return program.acceptState
}

// GetInitState returns the init state
func (program *Program) GetInitState() string {
	return program.initState
}

//NewProgram instantiates a new program with default values for init and accept states
func NewProgram() *Program {
	return &Program{"UNNAMED", make(map[string]State, maxStates), "q0", "qAccept", "q0"}
}

// Execute executes the program against a tape, and returns whether the input is accepted
func (program *Program) Execute(tape *Tape, flags ...string) bool {

	currState := false
	tapeState := false

	for _, v := range flags {
		if v == "CURRENT_STATE" {
			currState = true
			continue
		}
		if v == "TAPE_STATE" {
			tapeState = true
			continue
		}
	}

	for program.currState != program.acceptState {

		symbol := tape.Read()

		if currState {
			fmt.Println("DEBUG:", "current state:", program.currState)
		}

		if tapeState {
			fmt.Println(tape.GetRepresentation())
		}

		action, ok := program.getStateByName(program.currState).Mappings[symbol]
		if !ok {
			//INVALID STATE
			return false
		}

		nextState, instruction := action.ExtractInformation()
		tape.ExecuteTapeInstruction(instruction)

		program.currState = nextState
	}

	return true
}
