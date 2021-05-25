package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type playerData struct {
	name  string
	class model.characterClass
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
	fmt.Println("Your class is:", player.class.name)
	fmt.Println("  Strength:", player.class.strength)
	fmt.Println("  Agility:", player.class.agility)
	fmt.Println("  Charisma:", player.class.charisma)
}

func CharacterCreation(player *playerData) {
	for {
		err := CreateCharacter(player)
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
	if text == "y" || text == "Y" {
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

func CreateCharacter(player *playerData) error {

	fmt.Println()
	fmt.Println("Enter your character's name:")
	fmt.Print(inputMarker)
	text := GetPlayerInput()
	player.name = text

	fmt.Println()
	fmt.Println("Choose Your character's class:")
	for i := 0; i < len(classes); i++ {
		fmt.Println(" ", strconv.Itoa(i)+".", classes[i].name)
	}
	fmt.Print(inputMarker)

	text = GetPlayerInput()
	if _, err := strconv.Atoi(text); err == nil {

		index, _ := strconv.Atoi(text)
		if index >= 0 && index < len(classes) {
			player.class = classes[index]
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
