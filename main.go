package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"

	"github.com/RileyGabrielson/textAdventure/book1"
	"github.com/RileyGabrielson/textAdventure/model"
)

func main() {
	DisplayWelcome()
	player := playerData{}
	NewCharacter(&player)
	StartBook(&player, &book1.Book1)
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
	fmt.Println("  Intelligence:", player.class.Intelligence)
}

func AssignName(player *playerData) {
	fmt.Println()
	fmt.Println("Enter your character's name:")
	text := GetPlayerInput()
	player.name = text
}

func ClearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func StartBook(player *playerData, book *model.Book) {
	ClearScreen()
	fmt.Println()
	fmt.Println(book.Title)
	fmt.Println(book.Description)
	StartChapter(player, &book.Chapters[0])
}

func StartChapter(player *playerData, chapter *model.Chapter) {
	fmt.Println()
	fmt.Println(chapter.Title)
	fmt.Println(chapter.Description)
	StartChoice(player, &chapter.StartingChoice)
}

func StartChoice(player *playerData, choice *model.Choice) {
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
