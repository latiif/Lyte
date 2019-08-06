package main

import (
	"fmt"
	"os"
	"strings"

	"./ast"
	"./compiler"
	"./parse"
)

func main() {
	filename := os.Args[1]

	program, parsed := parse.Parse(filename)

	if parsed {
		fmt.Println("Program successfully parsed!")
	} else {
		fmt.Println("Invalid Lyte code\nHalting.")
		return
	}

	fmt.Print("Enter Tape Contents: ")
	var rawTape string
	fmt.Scanln(&rawTape)
	tape := ast.NewTape(255, strings.Split(rawTape, ""))
	accepted := program.Execute(tape)

	if accepted {
		fmt.Println("Accepted")
	} else {
		fmt.Println("Rejected")
	}

	fmt.Println("Tape: ", tape.GetRepresentation())

	//compiler.GoCompile(program)
}
