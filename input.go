package main

var inputMarker = "-> "

func GetPlayerInput() string {
	fmt.Print(inputMarker)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}

func GetPlayerInt() (int, error) {
	text := GetPlayerInput()
	if _, err := strconv.Atoi(text); err == nil {

		playerInt, _ := strconv.Atoi(text)
		return playerInt, nil

	} else {
		return -1, errors.New("invalid player input")
	}
}