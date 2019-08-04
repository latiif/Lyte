package ast

import (
	"strings"
)

type Tape struct {
	contents []string
	headPos  int32
}

const tapesize = 255

func NewTape(size uint32, initialContent []string) *Tape {

	mContents := make([]string, tapesize, tapesize)

	var actualStart = int32(tapesize / 2)

	for i := 0; i < len(initialContent); i++ {
		mContents[i+int(actualStart)] = initialContent[i]
	}

	tape := &Tape{mContents, actualStart}
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
		tape.contents[tape.headPos] = ""
	}

	tape.contents[tape.headPos] = str
}

func (tape *Tape) moveHead(direction string) {
	switch direction {
	case ">":
		tape.headPos++
	case "<":
		tape.headPos--
	case "-":
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
