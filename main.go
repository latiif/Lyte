package main

import (
	"fmt"

	"./ast"
	"./test"
)

func main() {

	p := test.CreateProgramByParse()

	tape := ast.NewTape(255, []string{"0", "1", "0"})

	ok := p.Execute(tape)

	fmt.Println(ok)
	fmt.Println(tape.GetRepresentation())

}
