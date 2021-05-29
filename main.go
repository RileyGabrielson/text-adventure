package main

import (
	"github.com/RileyGabrielson/text-adventure/book1"
	"github.com/RileyGabrielson/text-adventure/model"
	"github.com/RileyGabrielson/text-adventure/ui"
)

func main() {
	ui.DisplayWelcome()
	player := model.PlayerData{}
	ui.NewCharacter(&player)
	ui.StartBook(&player, &book1.Book1)
}


