package main

import "fmt"

type Mob struct {
	name  string
	hp    int
	lvl   int
	stm, str, int, def int
	xp    int
}

func (mob* Mob) show(x, y int) {
	lines := []string{
		mob.name,
		fmt.Sprintf("%d/%d", mob.hp, mob.stm*5),
		fmt.Sprintf("lvl: %d", mob.lvl),
		fmt.Sprintf("stm: %d", mob.stm),
		fmt.Sprintf("str: %d", mob.str),
		fmt.Sprintf("int: %d", mob.int),
		fmt.Sprintf("def: %d", mob.def),
		fmt.Sprintf("xp:  %d", mob.xp),
	}

	screen.PutStr(x, y, "╔════════╗")
	for i, l := range(lines) {
		screen.PutStr(x, y+i+1, fmt.Sprintf("║%-8s║", l))
	}
	screen.PutStr(x, y+len(lines)+1, "╚════════╝")
}

func (mob* Mob) fullhp() {
	mob.hp = mob.stm * STAMMOD
}
