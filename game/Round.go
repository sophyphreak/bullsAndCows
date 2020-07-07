package game

import "time"

// Round contains data for one round
type Round struct {
	Number         int
	Guess          int
	ExactMatches   int
	PartialMatches int
	TimeOfGuess    time.Time
}
