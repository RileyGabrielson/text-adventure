package main

import (
	"fmt"
	"strconv"

	"github.com/RileyGabrielson/text-adventure/book1"
	"github.com/RileyGabrielson/text-adventure/model"
	"github.com/RileyGabrielson/text-adventure/ui"

)

func main() {
	ui.DisplayWelcome()
	player := model.PlayerData{}
	ui.NewCharacter(&player)
	StartBook(&player, &book1.Book1)
}

func StartBook(player *model.PlayerData, book *model.Book) {
	ui.ClearScreen()
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

	index, err := ui.GetPlayerInt()
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
