package main

import (
	"fmt"

	"./test"
)

func main() {

	p := test.CreateProgram()
	ok, tape := p.Execute()
	fmt.Println(ok)
	fmt.Println(tape)
}
