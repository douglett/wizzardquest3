package main

import "fmt"
import "github.com/gdamore/tcell/v3"

var screen tcell.Screen

var player = Mob{
	name: "player",
}
var enemy = Mob{
	name: "slime",
}

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
	screen.SetTitle("Wizzard Battle Engine!")

	player.hp = player.stm * 5
	enemy.hp = enemy.stm * 5
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
			screen.PutStr(0, 0, fmt.Sprintf("Hello world %d", frame))
			player.show(10, 10)
			enemy.show(50, 10)
		screen.Show()
		frame++
	}
}
