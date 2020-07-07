package game

import "time"

// Round contains data for one round
type Round struct {
	Number         int       `json:"number"`
	Guess          int       `json:"guess"`
	ExactMatches   int       `json:"exactMatches"`
	PartialMatches int       `json:"partialMatches"`
	TimeOfGuess    time.Time `json:"timeOfGuess"`
}
