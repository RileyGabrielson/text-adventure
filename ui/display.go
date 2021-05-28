package ui

import (
	"fmt"
	"os/exec"
	"os"

	"github.com/RileyGabrielson/text-adventure/model"
)

func DisplayWelcome() {
	ClearScreen()
	fmt.Println("----------- Blarg -----------")
	fmt.Println("-- An adventure in Golang --")
	fmt.Println()
}

func DisplayCharacterDetails(player *model.PlayerData) {
	fmt.Println()
	fmt.Println("Your name is:", player.Name)
	fmt.Println("Your class is:", player.Class.Name)
	fmt.Println("  Strength:", player.Class.Strength)
	fmt.Println("  Agility: ", player.Class.Agility)
	fmt.Println("  Charisma:", player.Class.Charisma)
	fmt.Println("  Intelligence:", player.Class.Intelligence)
}

func ClearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}