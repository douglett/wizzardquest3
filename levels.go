package main
import "fmt"

var LevelStats = []Mob{
	Mob{},
	Mob{
		name: "player",
		lvl: 1,
		xp:  0,
		stm: 3,
		str: 3,
	},
	Mob{
		lvl: 2,
		xp:  20,
		stm: 5,
		str: 5,
	},
}

func levelup() {
	if player.lvl >= len(LevelStats) - 1 { return }
	next := LevelStats[player.lvl + 1]
	if player.xp < next.xp { return }

	// level up
	dialog.append("")
	dialog.append(fmt.Sprintf("You have reached level %d!", next.lvl))
	dialog.append(fmt.Sprintf("You gain: stm+%d, str+%d", next.stm - player.stm, next.str - player.str))
	
	player.lvl = next.lvl
	player.stm = next.stm
	player.str = next.str
	player.xp  = player.xp - Lvl2.xp

	waitspace()
}