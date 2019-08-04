package ast

import (
	"strings"
)

type Tape struct {
	contents []string
	headPos  int32
}

func NewTape(size uint32, initialContent []string) *Tape {

	mContents := make([]string, 255, 255)

	for i := 0; i < len(initialContent); i++ {
		mContents[i] = initialContent[i]
	}

	tape := &Tape{mContents, 0}
	return tape
}

func (tape *Tape) Read() string {

	symbol := tape.contents[tape.headPos]
	if symbol == "" {
		return "_"
	}

	return symbol

}

func (tape *Tape) write(str string) {

	if str == "_" {
		return
	}

	tape.contents[tape.headPos] = str
}

func (tape *Tape) moveHead(direction string) {
	switch direction {
	case ">":
		tape.headPos++
	case "<":
		tape.headPos--
	case "_":
		break
	}
}

func (tape *Tape) ExecuteTapeInstruction(instruction TapeInstruction) {
	tape.write(instruction.stringToWrite)
	tape.moveHead(instruction.headDirection)
}

func (tape *Tape) GetRepresentation() string {
	return strings.Join(tape.contents[:], "")
}
