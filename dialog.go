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
	if di.error != "" {
		screen.PutStr(x, y+len(di.lines)+1, di.error)
	}
}
