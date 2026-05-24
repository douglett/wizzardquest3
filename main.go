package main

import "fmt"
import "github.com/gdamore/tcell/v3"

var screen tcell.Screen
var dialog = Dialog{}
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
		// show
		dialog.lines = []string{
			"What will you do?",
			"  (a)ttack",
			"  (r)run",
		}
		show()

		// input
		switch s := input(); s {
			case "", "resize":
			case "quit", "r":
				break loop
			case "a":
				attack()
			// default:
			// 	dialog.error = "unknown command: " + s
		}
	}
}

func attack() {
	playerdmg := maxi(1, player.str - enemy.def)
	enemy.hp -= playerdmg
	dialog.lines = []string{
		fmt.Sprintf("You attack the %s for %d damage!", enemy.name, playerdmg),
	}

	loop1: for {
		show()
		switch input() {
			case " ":  break loop1
		}
	}

	enemydmg := maxi(1, enemy.str - player.def)
	player.hp -= enemydmg
	dialog.lines = append(dialog.lines, "")
	dialog.lines = append(dialog.lines, fmt.Sprintf("%s attacks you for %d damage!", enemy.name, enemydmg))

	loop2: for {
		show()
		switch input() {
			case " ":  break loop2
		}
	}
}

func maxi(a, b int) int {
	if (a > b) { return a }
	return b
}

func show() {
	screen.Clear()
		screen.PutStr(0, 0, "Wizzard Battle Engine!")
		player.show(1, 2)
		enemy.show(40, 2)
		dialog.show(0, 12)
	screen.Show()
}

func input() string {
	// get events
	event := <-screen.EventQ()
	switch event := event.(type) {
		case *tcell.EventResize:
			// width, height = event.Size()
			return "resize"
		case *tcell.EventKey:
			switch event.Key() {
				case tcell.KeyCtrlC, tcell.KeyESC:
					return "quit"
				case tcell.KeyRune:
					return event.Str()
			}
	}
	// unhandled
	return ""
}
