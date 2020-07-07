package game

import "time"

// Games contains all games, active and finished
var Games []*game

type game struct {
	id     int
	answer [4]int
	status string
	rounds []round
}

type round struct {
	number         int
	guess          int
	exactMatches   int
	partialMatches int
	timeOfGuess    time.Time
}
