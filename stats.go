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
	Mob{
		lvl: 5,
		xp:  250,
		stm: 11,
		str: 10,
		def: 4,
	},
	// eyeballed
	Mob{
		lvl: 6,
		xp:  350,
		stm: 12,
		str: 11,
		def: 5,
	},
	Mob{
		lvl: 7,
		xp:  450,
		stm: 12,
		str: 12,
		def: 6,
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
	name: "bandit",
	lvl: 4,
	xp:  25,
	stm: 6,
	str: 10,
	def: 2,
}
var MobDarkKnight = Mob{
	name: "d-knight",
	lvl: 5,
	xp:  40,
	stm: 6,
	str: 10,
	def: 5,
}
var MobDarkLord = Mob{
	name: "d-lord",
	lvl: 6,
	xp:  100,
	stm: 9,
	str: 14,
	def: 4,
}
