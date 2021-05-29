package model

type CharacterClass struct {
	Name         string

	Strength     int
	Dexterity	 int
	Charisma     int
	Intelligence int
	Constitution int
	Wisdom		 int
}

var barbarian = &CharacterClass{
	Name:         "Barbarian",
	Strength:     18,
	Dexterity:    13,
	Charisma:     12,
	Intelligence: 8,
	Constitution: 16,
	Wisdom: 	  13,
}

var archMage = &CharacterClass{
	Name:         "Arch Mage",
	Strength:     9,
	Dexterity:    11,
	Charisma:     12,
	Intelligence: 14,
	Constitution: 11,
	Wisdom: 	  14,
}



var Classes = []CharacterClass{
	*barbarian,
	*archMage,
	*dragonKeeper,
	*shadowHunter}
