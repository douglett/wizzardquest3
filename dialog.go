package main

// import "fmt"

type Dialog struct {
	lines []string
	error string
}

func (di *Dialog) show(x, y int) {
	for i, l := range(di.lines) {
		screen.PutStr(x, y+i, l)
	}
	// error message
	screen.PutStr(x, y+len(di.lines)+1, di.error)
}

func (di *Dialog) clear() {
	di.lines = []string{}
	di.error = ""
}

func (di *Dialog) append(str string) {
	di.lines = append(di.lines, str)
}
