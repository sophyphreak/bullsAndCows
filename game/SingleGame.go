package game

func GetGame(gameID int) (Game, error) {
	gameCopy, err := findGameByID(gameID)
	if err == nil {
		emptyArr := [4]int{}
		gameCopy.Answer = emptyArr
	}
	return gameCopy, err
}
