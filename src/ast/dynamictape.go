package ast

type Entry struct {
	Value string
	Right *Entry
	Left  *Entry
}

func NewEntry(value string) *Entry {
	entry := &Entry{value, nil, nil}
	return entry
}

func (entry *Entry) Write(newValue string) {
	if newValue == "_" {
		entry.Value = ""
	}
	entry.Value = newValue
}

func (entry *Entry) Read() string {
	symbol := entry.Value
	if symbol == "" {
		return "_"
	}
	return symbol
}

type DynamicTape struct {
	current *Entry
}

func NewDynamicTape() *DynamicTape {
	tape := &DynamicTape{NewEntry("")}
	return tape
}

func (tape *DynamicTape) MoveHead(direction string) {
	switch direction {
	case ">":
		if tape.current.Right == nil {
			tape.current.Right = NewEntry("")
		}
		tape.current.Right.Left = tape.current
		tape.current = tape.current.Right
	case "<":
		if tape.current.Left == nil {
			tape.current.Left = NewEntry("")
		}
		tape.current.Left.Right = tape.current
		tape.current = tape.current.Left
	case "-":
		break
	}
}

func (tape *DynamicTape) Write(val string) {
	tape.current.Write(val)
}

func (tape *DynamicTape) Read() string {
	return tape.current.Read()
}
