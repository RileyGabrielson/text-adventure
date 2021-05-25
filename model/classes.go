package model

type CharacterClass struct {
	Name     string
	Strength int
	Agility  int
	Charisma int
}

var barbarian = &CharacterClass{
	Name:     "Barbarian",
	Strength: 12,
	Agility:  5,
	Charisma: 5,
}

var archMage = &CharacterClass{
	Name:     "Arch Mage",
	Strength: 4,
	Agility:  8,
	Charisma: 14,
}

var Classes = []CharacterClass{*barbarian, *archMage}
