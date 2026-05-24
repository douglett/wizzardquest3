package main

// import "fmt"
import "github.com/gdamore/tcell/v3"

var screen tcell.Screen

var player = Mob{
	name: "player",
	stm: 3,
	str: 3,
}
var enemy = Mob{
	name: "slime",
	stm: 1,
	str: 1,
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
					case tcell.KeyRune:
						if event.Str() == "r" {
							break loop
						}
				}
		}

		// show
		screen.Clear()
			screen.PutStr(0, 0, "Wizzard Battle Engine!")
			player.show(1, 2)
			enemy.show(40, 2)

			lines := []string{
				"What will you do?",
				"  (a)ttack",
				"  (r)run",
			}
			for i, l := range(lines) {
				screen.PutStr(0, 12+i, l)
			}
		screen.Show()
	}
}
