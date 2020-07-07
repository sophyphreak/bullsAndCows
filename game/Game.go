package game

// Games contains all games, active and finished
var Games []*game

// Game contains a single game's data
type Game struct {
	ID     int
	Answer [4]int
	Status string
	Rounds []Round
}
