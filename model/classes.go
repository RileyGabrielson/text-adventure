package model

type CharacterClass struct {
	Name         string
	Strength     int
	Agility      int
	Charisma     int
	Intelligence int
}

var barbarian = &CharacterClass{
	Name:         "Barbarian",
	Strength:     18,
	Agility:      5,
	Charisma:     5,
	Intelligence: 2,
}

var archMage = &CharacterClass{
	Name:         "Arch Mage",
	Strength:     4,
	Agility:      8,
	Charisma:     14,
	Intelligence: 2,
}

var dragonKeeper = &CharacterClass{
	Name:         "Dragon Keeper",
	Strength:     9,
	Agility:      9,
	Charisma:     5,
	Intelligence: 7,
}

var shadowHunter = &CharacterClass{
	Name:         "Shadow Hunter",
	Strength:     6,
	Agility:      12,
	Charisma:     9,
	Intelligence: 8,
}

var Classes = []CharacterClass{
	*barbarian,
	*archMage,
	*dragonKeeper,
	*shadowHunter}
