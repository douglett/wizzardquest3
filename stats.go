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
		def: 0,
	},
	Mob{
		lvl: 2,
		xp:  20,
		stm: 5,
		str: 5,
		def: 1,
	},
	Mob{
		lvl: 3,
		xp:  70,
		stm: 8,
		str: 7,
		def: 2,
	},
	Mob{
		lvl: 4,
		xp:  150,
		stm: 9,
		str: 9,
		def: 3,
	},
}

// Mob types
var MobSlime = Mob{
	name: "slime",
	lvl: 1,
	xp:  5,
	stm: 1,
	str: 2,
	def: 0,
}
var MobBat = Mob{
	name: "bat",
	lvl: 2,
	xp:  10,
	stm: 3,
	str: 3,
	def: 1,
}
var MobBison = Mob{
	name: "bison",
	lvl: 3,
	xp:  15,
	stm: 7,
	str: 5,
	def: 2,
}
var MobBandit = Mob{
	name: "bison",
	lvl: 4,
	xp:  20,
	stm: 6,
	str: 10,
	def: 2,
}
