package compiler

import (
	"bytes"
	"os"
	"text/template"

	"../ast"
)

// GoCompile transpilers the ast representation into Go Code
func GoCompile(program *ast.Program) {
	states := program.GetStatesCopy()

	var buffer bytes.Buffer

	for _, v := range states {
		buffer.WriteString("\n" + v.Name + ":")
		for chRead, action := range v.Mappings {
			buffer.WriteString("\nif head() == \"" + chRead + "\" {\n")
			chWrite, dir := action.GetInstruction()
			nextState := action.GetNextState()

			if chWrite != "_" {
				buffer.WriteString("\twrite(\"" + chWrite + "\")\n")
			}

			switch dir {
			case ">":
				buffer.WriteString("\tmoveright()")
			case "<":
				buffer.WriteString("\tmoveleft()")
			}

			buffer.WriteString("\n\tgoto " + nextState)

			buffer.WriteString("\n}")
		}
	}

	tmpl, _ := template.ParseFiles("./compiler/go.lyte.txt")

	type info struct {
		Name        string
		Code        string
		AcceptState string
	}

	tmpl.Execute(os.Stdout, info{"Name", buffer.String(), program.GetAcceptState()})

}
