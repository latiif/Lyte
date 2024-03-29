package compiler

func getTemplate() string {
	return `package main

import (
    "fmt"
    "strings"
)

var TheTape *Tape

func main(){

    var line string
    fmt.Scanln(&line)

    TheTape = NewTape(255,strings.Split(line,""))

    execute()

	fmt.Println(TheTape.GetRepresentation())

}

func head() string {
	return TheTape.Read()
}

func write(symbol string) {
	TheTape.write(symbol)
}

func moveright(){
	TheTape.moveHead(">")
}

func moveleft(){
	TheTape.moveHead("<")
}

func execute(){
var head_ string
goto {{.InitState}}

    {{.Code}}
{{ if not .AcceptStateMentioned }}
	{{.AcceptState}}:
		return
{{ end }}
}

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



func (tape *Tape) GetRepresentation() string {
	representation := strings.Replace(strings.Join(tape.contents[:], ""), "_", "", -1)
	if representation == "" {
		return "{EMPTY}"
	} else {
		return representation
	}
}
`
}
