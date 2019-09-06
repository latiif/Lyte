package main

import (
	"fmt"
	"os"
	"strings"

	"./ast"
	"./compiler"
	"./parse"
)

// Debug tells whether the program in Debug mode
var Debug bool

func main() {

	if len(os.Args) > 1 {
		cliMode()
	} else {
		interactiveMode()
	}

}

func interactiveMode() {

	fmt.Println(`🥳 Interactive Mode
Enter the filename of your Lytecode:`)

	var filename string
	fmt.Scanln(&filename)

	program, parsed := parse.Parse(filename)

	if parsed {
		fmt.Println("✔️ Program successfully parsed!")
	} else {
		fmt.Println("✋ Invalid Lyte code\nHalting.")
		return
	}

	fmt.Print("Enter Tape Contents: ")
	var rawTape string
	fmt.Scanln(&rawTape)
	tape := ast.NewTape(255, strings.Split(rawTape, ""))
	accepted := program.Execute(tape)

	if accepted {
		fmt.Println("✔️ Accepted")
	} else {
		fmt.Println("✋ Rejected")
	}

	fmt.Println("Tape: ", tape.GetRepresentation())

}

func cliMode() {
	filename := os.Args[1]

	program, parsed := parse.Parse(filename)

	if !parsed {
		fmt.Println("Invalid Lyte code!\nAborting.")
		return
	}

	os.Stdout.Write(compiler.GoCompile(program))
}
