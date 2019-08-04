package ast

type Action struct {
	nextState       string
	tapeInstruction TapeInstruction
}

func NewAction(nextState string, tapeInstruction TapeInstruction) Action {
	return Action{nextState, tapeInstruction}
}

func (action *Action) ExtractInformation() (string, TapeInstruction) {
	return action.nextState, action.tapeInstruction
}
