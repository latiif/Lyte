package compiler

import (
	"bytes"
	"text/template"

	"../ast"
)

// GoCompile transpilers the ast representation into Go Code
func GoCompile(program *ast.Program) []byte {
	states := program.GetStatesCopy()

	var buffer bytes.Buffer

	for _, v := range states {
		buffer.WriteString(parseState(v))
	}

	tmpl, _ := template.ParseFiles("./compiler/go.lyte.tpl")

	type info struct {
		Name        string
		Code        string
		InitState   string
		AcceptState string
	}

	var compiled bytes.Buffer
	tmpl.Execute(&compiled, info{
		"Name",
		buffer.String(),
		program.GetInitState(),
		program.GetAcceptState()})

	return compiled.Bytes()

}

func parseState(v ast.State) string {

	var buffer bytes.Buffer

	buffer.WriteString("\n" + v.Name + ":")
	buffer.WriteString("\nhead_ = head()")
	for chRead, action := range v.Mappings {
		buffer.WriteString("\nif head_ == \"" + chRead + "\" {\n")
		chWrite, dir := action.GetInstruction()
		nextState := action.GetNextState()

		buffer.WriteString("\twrite(\"" + chWrite + "\")\n")

		switch dir {
		case ">":
			buffer.WriteString("\tmoveright()")
		case "<":
			buffer.WriteString("\tmoveleft()")
		}

		//buffer.WriteString("\n\t" + `fmt.Println("Going to:","` + nextState + `")`)
		//buffer.WriteString("\n\t" + `fmt.Println(TheTape.GetRepresentation())`)

		buffer.WriteString("\n\tgoto " + nextState)

		buffer.WriteString("\n}")
	}

	buffer.WriteString("\nreturn")

	return buffer.String()
}
