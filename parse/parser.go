package parse

import (
	"io/ioutil"
	"strings"

	"../ast"
)

func Parse(filename string) (*ast.Program, bool) {
	b, _ := ioutil.ReadFile(filename)
	contents := strings.TrimSpace(string(b))

	lines := strings.Split(contents, "\n")

	program := ast.NewProgram("q0", "qAccept")

	for i := 0; i < len(lines); i += 2 {
		if lines[i] == "" {
			i--
			continue
		}

		fstRow := strings.Split(lines[i], ",")
		sndRow := strings.Split(lines[i+1], ",")

		if len(fstRow) != 2 || len(sndRow) != 3 {
			return nil, false
		}

		program.AddRule(fstRow[0], sndRow[0], fstRow[1], sndRow[1], sndRow[2])
	}

	return program, true
}
