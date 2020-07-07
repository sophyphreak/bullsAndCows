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

	check := ListRounds(0)
	if game1.Rounds[0] != round1 {
		t.Errorf("Expected round[0] to be %v but got %v", game1.Rounds[0], check)
	}

}
