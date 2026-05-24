package main

import "fmt"
import "github.com/gdamore/tcell/v3"

var screen tcell.Screen

func main() {
	var err error
	screen, err = tcell.NewScreen()
	defer screen.Fini()
	if err != nil {
		panic(err)
	}
	if err := screen.Init(); err != nil {
		panic(err)
	}

	frame := 0
	loop: for {
		// get events
		event := <-screen.EventQ()
		switch event := event.(type) {
			// case *tcell.EventResize:
			// 	width, height = event.Size()
			case *tcell.EventKey:
				switch event.Key() {
					case tcell.KeyCtrlC, tcell.KeyESC:
						break loop
				}
		}

		// show
		screen.Clear()
		screen.PutStr(10, 10, fmt.Sprintf("Hello world %d", frame))
		screen.Show()
		frame++
	}
}