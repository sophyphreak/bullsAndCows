package game

func GetGame(gameID int) (Game, error) {
	indx, err := findGameByID(gameID)
	gameCopy := *Games[indx]
	if err == nil {
		if gameCopy.Status == "IN PROGRESS" {
			gameCopy.Answer = [4]int{}
		}
	}
	return gameCopy, err
}
