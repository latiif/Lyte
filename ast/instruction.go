package ast

type TapeInstruction struct {
	stringToWrite string
	headDirection string
}

func NewTapeInstruction(stringToWrite, headDirection string) TapeInstruction {
	return TapeInstruction{
		stringToWrite: stringToWrite,
		headDirection: headDirection}
}
