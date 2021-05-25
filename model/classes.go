package model

type characterClass struct {
	name     string
	strength int
	agility  int
	charisma int
}

var barbarian = &characterClass{
	name:     "Barbarian",
	strength: 12,
	agility:  5,
	charisma: 5,
}

var archMage = &characterClass{
	name:     "Arch Mage",
	strength: 4,
	agility:  8,
	charisma: 14,
}

var classes = []characterClass{*barbarian, *archMage}
