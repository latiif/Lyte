package compiler

import (
	"fmt"

	"../ast"
)

// GoCompile transpilers the ast representation into Go Code
func GoCompile(program *ast.Program) {
	states := program.GetStatesCopy()

	for _, v := range states {
		fmt.Println(v.Name, ":")
		for chRead, action := range v.Mappings {
			fmt.Print("if head() == \"", chRead, "\" {\n")
			chWrite, dir := action.GetInstruction()
			nextState := action.GetNextState()


			if chWrite != "_" {
				fmt.Print("\twrite(\"",chWrite,"\")\n")
			}

			switch dir {
			case ">":
				fmt.Println("\tmoveright()")
			case "<":
				fmt.Println("\tmoveleft()")
			}

			fmt.Println("\tgoto ", nextState)

			fmt.Println("}")
		}
	}

}
