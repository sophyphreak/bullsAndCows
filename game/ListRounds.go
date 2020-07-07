package game

// ListRounds returns a list of all rounds in a game
func ListRounds(id int) []Round {

	index, err := findGameByID(id)
	foundGame := *(Games[index])
	if err != nil || index == -1 {
		panic(err)
	}

	return foundGame.Rounds
}
