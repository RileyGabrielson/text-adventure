package main

func AssignClass(player *playerData) error {

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
			player.class = model.Classes[index]
			return nil
		} else {
			return errors.New("invalid index")
		}
	}
}

func NewCharacter(player *playerData) {
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
		CharacterCreation(player)
	}

}