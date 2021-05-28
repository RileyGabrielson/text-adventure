package ui

import (
	"fmt"
	"strconv"
	"errors"
	"github.com/RileyGabrielson/text-adventure/model"
	"github.com/RileyGabrielson/text-adventure/commands"
)

func AssignClass(player *model.PlayerData) error {

	fmt.Println()
	fmt.Println("Choose Your character's class:")
	for i := 0; i < len(model.Classes); i++ {
		fmt.Println(" ", strconv.Itoa(i)+".", model.Classes[i].Name)
	}

	index, err := GetPlayerInt()
	if err != nil {
		return err
	} else {
		if index >= 0 && index < len(model.Classes) {
			player.Class = model.Classes[index]
			return nil
		} else {
			return errors.New("invalid index")
		}
	}
}

func NewCharacter(player *model.PlayerData) {
	AssignName(player)

	for {
		err := AssignClass(player)
		if err == nil {
			break
		} else {
			fmt.Println(err.Error())
		}
	}

	DisplayCharacterDetails(player)
	fmt.Println()
	fmt.Println("Continue? [Y/N]")
	text := GetPlayerInput()
	if _, containsKey := commands.YesCommands[text]; containsKey {
		return
	} else {
		ClearScreen()
		NewCharacter(player)
	}

}

func AssignName(player *model.PlayerData) {
	fmt.Println()
	fmt.Println("Enter your character's name:")
	text := GetPlayerInput()
	player.Name = text
}