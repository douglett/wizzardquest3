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
var enemy = Mob{}

var Slime = Mob{
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

	selectenemy()
}

func selectenemy() {
	for {
		// show
		dialog.lines = []string{
			"Select enemy:",
			"  (1)   Slime",
			"  (esc) Quit",
		}
		show()

		// input
		switch s := input(); s {
			case "", "resize":
			case "quit":
				return
			case "1":
				r := battle(1)
				if r == -1 { return }
		}
	}
}

func battle(enemynum int) (result int) {
	switch enemynum {
		case 1:  enemy = Slime
	}
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
		switch input() {
			case "quit":
				return -1
			case "r":
				break loop
			case "a":
				r := attack()
				switch r {
					case -1:    return -1
					case 1, 2:  break loop
				}
		}
	}

	// cleanup
	if enemy.hp <= 0 { result = 1 }
	enemy = Mob{}
	return
}

func attack() int {
	// player attack
	playerdmg := maxi(1, player.str - enemy.def)
	enemy.hp -= playerdmg
	dialog.lines = []string{
		fmt.Sprintf("You attack the %s for %d damage!", enemy.name, playerdmg),
	}
	if enemy.hp <= 0 {
		dialog.lines = append(dialog.lines, "")
		dialog.lines = append(dialog.lines, fmt.Sprintf("The %s has been defeated.", enemy.name))
	}
	// wait for space
	loop1: for {
		show()
		switch input() {
			case "quit":  return -1
			case " ":     break loop1
		}
	}
	if enemy.hp <= 0 { return 1 }

	// enemy attack
	enemydmg := maxi(1, enemy.str - player.def)
	player.hp -= enemydmg
	dialog.lines = append(dialog.lines, "")
	dialog.lines = append(dialog.lines, fmt.Sprintf("%s attacks you for %d damage!", enemy.name, enemydmg))
	if player.hp <= 0 {
		dialog.lines = append(dialog.lines, "")
		dialog.lines = append(dialog.lines, "You have been defeated.")
	}
	// wait for space
	loop2: for {
		show()
		switch input() {
			case " ":  break loop2
		}
	}
	if player.hp <= 0 { return 2 }
	return 0
}

func maxi(a, b int) int {
	if (a > b) { return a }
	return b
}

func show() {
	screen.Clear()
		screen.PutStr(0, 0, "Wizzard Battle Engine!")
		player.show(1, 2)
		if enemy.name != "" { enemy.show(40, 2) }
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
