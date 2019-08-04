package ast

import "strconv"

const maxStates = 255

type Program struct {
	name        string
	states      map[string]State
	initState   string
	acceptState string
	currState   string
	tape        Tape
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

func (program *Program) AddRule(frState, toState, frChar, toChar string, dir string) {
	instruction := NewTapeInstruction(toChar, dir)
	action := NewAction(toState, instruction)
	state := program.getStateByName(frState)
	state.Mappings[frChar] = action
}

func NewProgram(init, accept string, tape Tape) Program {
	return Program{"", make(map[string]State, maxStates), init, accept, init, tape}
}

func (program *Program) Display() string {
	return strconv.Itoa(len(program.states["q1"].Mappings))
}

func (program *Program) Execute() (bool, string) {
	for program.currState != program.acceptState {
		symbol := program.tape.Read()
		action, ok := program.getStateByName(program.currState).Mappings[symbol]
		if !ok {
			//INVALID STATE
			return false, program.tape.GetRepresentation()
		}
		nextState, instruction := action.ExtractInformation()
		program.tape.ExecuteTapeInstruction(instruction)

		program.currState = nextState
	}

	return true, program.tape.GetRepresentation()
}
