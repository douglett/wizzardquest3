package main

import "fmt"

type Mob struct {
	name  string
	hp    int
	stm, str, int, def int
}

func (mob* Mob) show(x, y int) {
	lines := []string{
		mob.name,
		fmt.Sprintf("%d/%d", mob.hp, mob.stm*5),
		fmt.Sprintf("stm: %d", mob.stm),
		fmt.Sprintf("str: %d", mob.str),
		fmt.Sprintf("int: %d", mob.int),
		fmt.Sprintf("def: %d", mob.def),
	}

	// lines = append(lines, )

	// screen.PutStr(x, y, fmt.Sprintf("Hello world %d", frame))
	// screen.PutStr(x, y, lines[0])
	for i, line := range(lines) {
		screen.PutStr(x, y+i, line)
	}
}
