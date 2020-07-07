package game

// ListGames returns a slice of all game ids
func ListGames() []Game {

	var allGames []Game

	for i := 0; i < len(Games); i++ {
		game, err := GetGame(i)
		Check(err)
		allGames = append(allGames, game)
	}
	return allGames
}
