package compiler

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"text/template"

	"../ast"
)

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

// GoCompile transpilers the ast representation into Go Code
func GoCompile(program *ast.Program) []byte {
	states := program.GetStatesCopy()

	var buffer bytes.Buffer

	isAcceptingState := false

	for _, v := range states {
		go buffer.WriteString(parseState(v))
		isAcceptingState = isAcceptingState || (program.GetAcceptState() == v.Name)
	}

	tmpl, _ := template.ParseFiles("./compiler/go.lyte.tpl")

	type info struct {
		Name                 string
		Code                 string
		InitState            string
		AcceptState          string
		AcceptStateMentioned bool
	}

	var compiled bytes.Buffer
	tmpl.Execute(&compiled, info{
		"Name",
		buffer.String(),
		program.GetInitState(),
		program.GetAcceptState(),
		isAcceptingState,
	})

	if commandExists("go") {

		os.Mkdir(fmt.Sprintf("./%s", program.GetName()), os.ModePerm)
		ioutil.WriteFile(fmt.Sprintf("./%s/main.go", program.GetName()), compiled.Bytes(), 0644)

		cmd := exec.Command("go", "build")
		cmd.Dir = fmt.Sprintf("./%s", program.GetName())
		cmd.Run()
	}

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
