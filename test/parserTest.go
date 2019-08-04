package test

import (
	"../ast"
	"../parse"
)

func CreateProgramByParse() *ast.Program {
	program, _ := parse.Parse("decimaltobinary.lyte")
	return program
}
