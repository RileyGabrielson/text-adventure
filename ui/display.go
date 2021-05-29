package ui

import (
	"fmt"
	"os/exec"
	"os"
	"strconv"

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

func StartBook(player *model.PlayerData, book *model.Book) {
	ClearScreen()
	fmt.Println()
	fmt.Println(book.Title)
	fmt.Println(book.Description)
	StartChapter(player, &book.Chapters[0])
}

func StartChapter(player *model.PlayerData, chapter *model.Chapter) {
	fmt.Println()
	fmt.Println(chapter.Title)
	fmt.Println(chapter.Description)
	StartChoice(player, &chapter.StartingChoice)
}

func StartChoice(player *model.PlayerData, choice *model.Choice) {
	fmt.Println()
	fmt.Println(choice.Location)
	fmt.Println(choice.Description)

	if len(choice.Options) == 0 {
		return
	}

	for i := 0; i < len(choice.Options); i++ {
		fmt.Println(" ", strconv.Itoa(i)+".", choice.Options[i].Decision)
	}

	index, err := GetPlayerInt()
	if err != nil {
		fmt.Println()
		fmt.Println(err.Error())
		StartChoice(player, choice)
	} else {
		if index >= 0 && index < len(choice.Options) {
			StartChoice(player, &choice.Options[index])
			//New Comment again
		} else {
			fmt.Println()
			fmt.Println("Invalid Index")
			StartChoice(player, choice)
		}
	}

}