package game

// ListGames returns a slice of all game ids
func ListGames() []Game {

	var allGames []Game

	for i := 0; i < len(Games); i++ {
		game, err := GetGame(i)
		if err != nil {
			panic(err)
		}
		allGames = append(allGames, game)
	}
	return allGames
}
