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
			fmt.Println("if head() == \"", chRead, "\" {")
			chWrite, dir := action.GetInstruction()
			nextState := action.GetNextState()

			switch chWrite {
			case "_":
				continue
			default:
				fmt.Println("\twrite(\"", chWrite, "\")")
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
