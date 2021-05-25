package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/RileyGabrielson/textAdventure/model"
)

type playerData struct {
	name  string
	class model.CharacterClass
}

var inputMarker = "-> "

func main() {
	DisplayWelcome()
	player := playerData{}
	CharacterCreation(&player)

}

func DisplayWelcome() {
	ClearScreen()
	fmt.Println("----------- Blarg -----------")
	fmt.Println("-- An adventure in Golang --")
	fmt.Println()
}

func DisplayCharacterDetails(player *playerData) {
	fmt.Println()
	fmt.Println("Your name is:", player.name)
	fmt.Println("Your class is:", player.class.Name)
	fmt.Println("  Strength:", player.class.Strength)
	fmt.Println("  Agility: ", player.class.Agility)
	fmt.Println("  Charisma:", player.class.Charisma)
}

func CharacterCreation(player *playerData) {
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
	if text == "y" || text == "Y" || text == "yes" ||  {
		return
	} else {
		ClearScreen()
		CharacterCreation(player)
	}

}

func GetPlayerInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}

func AssignName(player *playerData) {
	fmt.Println()
	fmt.Println("Enter your character's name:")
	fmt.Print(inputMarker)
	text := GetPlayerInput()
	player.name = text
}

func AssignClass(player *playerData) error {

	fmt.Println()
	fmt.Println("Choose Your character's class:")
	for i := 0; i < len(model.Classes); i++ {
		fmt.Println(" ", strconv.Itoa(i)+".", model.Classes[i].Name)
	}
	fmt.Print(inputMarker)

	text := GetPlayerInput()
	if _, err := strconv.Atoi(text); err == nil {

		index, _ := strconv.Atoi(text)
		if index >= 0 && index < len(model.Classes) {
			player.class = model.Classes[index]
			return nil
		} else {
			return errors.New("invalid index")
		}

	} else {
		return errors.New("invalid player input")
	}

}

func ClearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
