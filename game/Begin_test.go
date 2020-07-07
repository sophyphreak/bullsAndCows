package game

import "testing"

func TestBegin(t *testing.T) {
	Begin()
	var newGame Game
	if len(Games) == 0 {
		t.Errorf("Expected new game in Games, but Games is empty")
		newGame = Game{}
	} else {
		newGame = *Games[len(Games)-1]
	}
	if newGame.ID <= -1 {
		t.Errorf("Expected id is greater than -1 but instead received %d", newGame.ID)
	}
	if newGame.Status != "IN PROGRESS" {
		t.Errorf("Expected newGame.status to be IN PROGRESS but instead found %s", newGame.Status)
	}
}

func TestGenRandArray(t *testing.T) {
	answers := GenRandArray()
	freqCount := make(map[int]int)
	for _, n := range answers {
		freqCount[n]++
		if freqCount[n] >= 2 {
			t.Errorf("Expected [4]int with uniqe values got %v", answers)
			break
		}
	}
}
