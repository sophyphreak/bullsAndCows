package game

func GetGame(gameID int) (Game, error) {
	gameCopy, err := findGameByID(gameID)
	if err == nil {
		if gameCopy.Status = "inProgress" {
			emptyArr := [4]int{}
			gameCopy.Answer = emptyArr
		}
	}
	return gameCopy, err
}
