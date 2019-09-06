package ast

type State struct {
	Name     string
	Mappings map[string]Action
}

func NewState(name string) State {
	state := State{name, make(map[string]Action, 255)}
	return state
}
