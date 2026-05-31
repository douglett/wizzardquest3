package main

// stats
const STAMMOD = 5

// player levels
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

// Mob types
var MobSlime = Mob{
	name: "slime",
	lvl: 1,
	stm: 1,
	str: 2,
	xp:  5,
}
