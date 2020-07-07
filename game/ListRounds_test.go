package game

import (
	"testing"
	"time"
)

func TestListRounds(t *testing.T) {
	round1 := Round{
		1,
		1,
		3,
		1,
		time.Now(),
	}
	var round []Round
	round = append(round, round1)
	game1 := Game{
		0,
		[4]int{1, 2, 3, 4},
		"IN PROGRESS",
		round,
	}
	Games = append(Games, &game1)
	foundRounds := ListRounds(0)
	if foundRounds[0] != round1 {
		t.Errorf("Expected round1 to be %v but got %v", round1, foundRounds[0])
	}
	Games = Games[:0]
}
