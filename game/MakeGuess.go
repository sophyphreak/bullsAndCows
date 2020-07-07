package game

// MakeGuess handles making the guess
func MakeGuess(id int, g string) (Round, error) {
	game, err := findGameByID(id)
	if err != nil {
		return Round{}, err
	}

}
