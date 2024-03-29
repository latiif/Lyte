package test

import (
	"../ast"
)

func CreateProgram() *ast.Program {
	program := ast.NewProgram("q0", "qAccept")
	program.AddRule("q0", "q0", "0", "0", ">")
	program.AddRule("q0", "q1", "1", "1", ">")
	program.AddRule("q1", "q2", "0", "0", ">")
	program.AddRule("q1", "q0", "1", "1", ">")
	program.AddRule("q2", "q1", "0", "0", ">")
	program.AddRule("q2", "q2", "1", "1", ">")

	program.AddRule("q0", "qAccept", "_", "_", "-")

	return program
}
