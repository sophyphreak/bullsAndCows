package game

import (
	"strconv"
	"strings"
	"time"
)

// MakeGuess handles making the guess
func MakeGuess(id int, gu string) (Round, error) {
	ga, err := findGameByID(id)
	if err != nil {
		return Round{}, err
	}

	game := *ga

	guess := strings.Split(gu, "")
	round := Round{
		len(game.Rounds) + 1,
		len(game.Rounds) + 1,
		0,
		0,
		time.Now(),
	}
	for _, n := range game.Answer {
		for _, m := range guess {
			if strconv.Itoa(n) == m {
				round.PartialMatches++
			}
		}
	}
	for i, n := range game.Answer {
		if guess[i] == strconv.Itoa(n) {
			round.ExactMatches++
			round.PartialMatches--
		}
	}
	game.Rounds = append(game.Rounds, round)
	if round.ExactMatches == 4 {
		game.Status = "FINISHED"
	}
	return round, nil
}
