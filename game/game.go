package game

// Games contains all games, active and finished
var Games []*game

type game struct {
	id     int
	answer [4]int
	status string
	rounds []Round
}
