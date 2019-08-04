package parse

import (
	"io/ioutil"
	"strings"

	"../ast"
)

// Parse Given a file name returns a parsed program in AST form
func Parse(filename string) (*ast.Program, bool) {
	b, _ := ioutil.ReadFile(filename)
	contents := strings.TrimSpace(string(b))

	lines := strings.Split(contents, "\n")

	program := ast.NewProgram()

	for i := 0; i < len(lines); i++ {
		lines[i] = strings.Replace(lines[i], " ", "", -1)
		if lines[i] == "" {
			continue
		} else if strings.HasPrefix(lines[i], "//") {
			continue
		} else if strings.Contains(lines[i], ":") {

			entries := strings.Split(lines[i], ":")

			switch entries[0] {
			case "name":
				program.AddName(entries[1])
			case "init":
				program.AddInitState(entries[1])
			case "accept":
				program.AddAcceptState(entries[1])
			}

			continue

		} else {
			fstRow := strings.Split(lines[i], ",")
			sndRow := strings.Split(lines[i+1], ",")

			if len(fstRow) < 2 || len(sndRow) < 3 {
				return nil, false
			}

			program.AddRule(fstRow[0], sndRow[0], fstRow[1], sndRow[1], sndRow[2])
			i++
		}

	}

	return program, true
}
