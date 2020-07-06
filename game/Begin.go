package game

import (
	"math/rand"
	"time"
)

// Begin handles beginning a new game, returns id of that new game
func Begin() int {
	var newGame game
	rand.Seed(time.Now().UTC().UnixNano())
	// create the four unique, random digits
	answer := [4]int{0, 0, 0, 0}
	exists := make(map[int]bool)
	for i := range answer {
		n = rand.Intn(10)

	}
	// sets the status
	// adds it to the global slice
	// use make() to create empty rounds slice
	// returns id of new game
	return -1
}
