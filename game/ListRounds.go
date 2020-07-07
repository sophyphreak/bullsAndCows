package game

// ListRounds returns a list of all rounds in a game
func ListRounds(id int) []Round {

	foundGame, err := findGameByID(id)
	if err != nil {
		panic(err)
	}

	return foundGame.Rounds
}
