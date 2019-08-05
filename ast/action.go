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

func (action *Action) GetNextState() string {
	return action.nextState
}

func (action *Action) GetInstruction() (string, string) {
	return action.tapeInstruction.GetInstruction()
}
