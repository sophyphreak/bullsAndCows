package game

import "testing"

func TestBegin(t *testing.T) {
	Begin()
	var newGame game
	if len(Games) == 0 {
		t.Errorf("Expected new game in Games, but Games is empty")
		newGame = game{}
	} else {
		newGame = *Games[len(Games)-1]
	}
	if newGame.id <= -1 {
		t.Errorf("Expected id is greater than -1 but instead received %d", newGame.id)
	}
	if newGame.status != "IN PROGRESS" {
		t.Errorf("Expected newGame.status to be IN PROGRESS but instead found %s", newGame.status)
	}
	freqCount := make(map[int]int)
	for _, n := range newGame.answer {
		freqCount[n]++
		if freqCount[n] >= 2 {
			t.Errorf("Expected all numbers in newGame.answer to be unique but instead received %v", newGame.answer)
			break
		}
	}
}
