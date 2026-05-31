package main

import "fmt"
import "github.com/gdamore/tcell/v3"

var screen tcell.Screen
var dialog = Dialog{}
var player = Mob{}
var enemy = Mob{}

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

	player = LevelStats[6]
	player.fullhp()

	// temp
	// player.hp = 1
	player.name = "player"
	player.xp = 0

	selectenemy()
}

// === game ===

func recoverquit() {
	if r := recover(); r != nil {
		if r == "quit" {
			fmt.Println("game quit")
		} else {
			panic(r)
		}
	}
}

func selectenemy() (result int) {
	const REST_TURNS = 3
	rest := REST_TURNS

	for {
		// show
		dialog.lines = []string{
			fmt.Sprintf("Rest available in %d turns.", rest),
			"Select enemy:",
			"  (1)   Slime",
			"  (2)   Bat",
			"  (3)   Bison",
			"  (4)   Bandit",
			"  (5)   DarkKnight",
			"  (6)   DarkLord",
			"  ... ",
			"  (r)   Rest",
			"  (esc) Quit",
		}
		show()

		// input
		switch input() {
			case "1":
				enemy = MobSlime
				result = battle()
			case "2":
				enemy = MobBat
				result = battle()
			case "3":
				enemy = MobBison
				result = battle()
			case "4":
				enemy = MobBandit
				result = battle()
			case "5":
				enemy = MobDarkKnight
				result = battle()
			case "6":
				enemy = MobDarkLord
				result = battle()
				if result == 1 { result = 3 }
			case "r":
				dialog.append("")
				if rest > 0 {
					dialog.append("You can't rest again this turn.")
				} else {
					dialog.append("You rest for a while. You recover 20 hp.")
					player.addhp(20)
					rest = REST_TURNS
				}
				waitspace()
		}

		if result == 1 { rest = maxi(0, rest-1) }
		if result == 2 { return }
		if result == 3 { youwin(); return }
		result = 0
	}
}

func battle() (result int) {
	// begin at full enemy hp
	enemy.fullhp()

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
				result = attack()
				if result != 0 { break loop }
		}
	}

	// level up?
	if result == 1 {
		player.xp += enemy.xp
		dialog.append("")
		dialog.append(fmt.Sprintf("You gained %d XP.", enemy.xp))
		waitspace()
		levelup()
	}

	// cleanup
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

func levelup() {
	if player.lvl >= len(LevelStats) - 1 { return }
	next := LevelStats[player.lvl + 1]
	if player.xp < next.xp { return }

	// level up
	dialog.append("")
	dialog.append(fmt.Sprintf("You have reached level %d!", next.lvl))
	dialog.append(fmt.Sprintf("You gain: stm+%d, str+%d, def+%d", 
		next.stm - player.stm, 
		next.str - player.str,
		next.def - player.def))
	
	player.lvl = next.lvl
	player.stm = next.stm
	player.str = next.str
	player.xp  = player.xp - next.xp

	waitspace()
}

func youwin() {
	dialog.lines = []string{
		"You have defeated the Dark Lord!",
		"",
		"Peace reigns accross the land. Well done hero!",
	}

	waitspace()
}

// === screen and input ===

func show() {
	screen.Clear()
		screen.PutStr(0, 0, "Wizzard Battle Engine!")
		player.show(1, 2)
		if enemy.name != "" { enemy.show(40, 2) }
		dialog.show(0, 14)
	screen.Show()
}

func waitspace() int {
	for {
		show()
		if input() == " " { return 0 }
	}
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

// === helpers ===

func maxi(a, b int) int {
	if (a > b) { return a }
	return b
}

func mini(a, b int) int {
	if (a < b) { return a }
	return b
}
