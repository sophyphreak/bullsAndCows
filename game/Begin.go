package game

import (
	"math/rand"
	"time"
)

// Begin handles beginning a new game, returns id of that new game
func Begin() int {
	var newGame game
	//set id
	newGame.id = 0
	// create the four unique, random digits
	rand.Seed(time.Now().UTC().UnixNano())
	newGame.answer = GenRandArray()
	//set the status
	newGame.status = "inProgress"
	// use make() to create empty rounds slice
	newGame.rounds = make([]round, 8)
	// adds it to the global slice
	Games = append(Games, &newGame)
	// returns id of new game
	return newGame.id
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
