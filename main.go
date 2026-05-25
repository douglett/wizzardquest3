package main

import "fmt"
import "github.com/gdamore/tcell/v3"

const STAMMOD = 5
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
	str: 2,
}

func main() {
	var err error
	screen, err = tcell.NewScreen()
	defer recoverquit()
	defer screen.Fini()
	if err != nil {
		panic(err)
	}
	if err := screen.Init(); err != nil {
		panic(err)
	}
	screen.SetTitle("Wizzard Battle Engine!")

	player.hp = player.stm * STAMMOD
	// player.hp = 1

	selectenemy()
}

func recoverquit() {
	if r := recover(); r != nil {
		if r == "quit" {
			fmt.Println("game quit")
		} else {
			panic(r)
		}
	}
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
		switch input() {
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
	enemy.hp = enemy.stm * STAMMOD

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
			case "r":
				break loop
			case "a":
				if attack() != 0 { break loop }
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
	waitspace()
	
	// check for loss
	if enemy.hp <= 0 {
		dialog.lines = append(dialog.lines, "")
		dialog.lines = append(dialog.lines, fmt.Sprintf("The %s has been defeated.", enemy.name))
		waitspace()
		return 1
	}

	// enemy attack
	enemydmg := maxi(1, enemy.str - player.def)
	player.hp -= enemydmg
	dialog.lines = append(dialog.lines, "")
	dialog.lines = append(dialog.lines, fmt.Sprintf("%s attacks you for %d damage!", enemy.name, enemydmg))
	waitspace()

	// check for loss
	if player.hp <= 0 {
		dialog.lines = append(dialog.lines, "")
		dialog.lines = append(dialog.lines, "You have been defeated.")
		waitspace()
		return 2
	}

	// no loss
	return 0
}

func waitspace() int {
	for {
		show()
		if input() == " " { return 0 }
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
					// return "quit"
					panic("quit")
				case tcell.KeyRune:
					return event.Str()
			}
	}
	// unhandled
	return ""
}
