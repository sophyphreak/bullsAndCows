package game

import (
	"testing"
	"time"
)

func TestListGames(t *testing.T) {
	var g *Games
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

	check := ListGames(g)
	if g.Game[0] != game1 {
		t.Errorf("Expected Game[0] to be %v but got %v", g.Game[0], check)
	}

}
