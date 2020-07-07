package game

import (
	"math/rand"
	"time"
)

// Begin handles beginning a new game, returns id of that new game
func Begin() int {
	var newGame Game
	//set id
	newGame.ID = 0
	// create the four unique, random digits
	rand.Seed(time.Now().UTC().UnixNano())
	newGame.Answer = GenRandArray()
	//set the status
	newGame.Status = "inProgress"
	// use make() to create empty rounds slice
	newGame.Rounds = make([]Round, 1)
	// adds it to the global slice
	Games = append(Games, &newGame)
	// returns id of new game
	return newGame.ID
}

func GenRandArray() [4]int {
	answers := [4]int{}
	answerMap := make(map[int]bool)
	for indx, _ := range answers {
		n := rand.Intn(10)
		for answerMap[n] == true {
			n = rand.Intn(10)
		}
		answerMap[n] = true
		answers[indx] = n
	}
	return answers
}
