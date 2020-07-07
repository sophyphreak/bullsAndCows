package game

import (
	"errors"
	"strconv"
)

// Games contains all games, active and finished
var Games []*Game

// Game contains a single game's data
type Game struct {
	ID     int
	Answer [4]int
	Status string
	Rounds []Round
}

func findGameByID(id int) (*Game, error) {
	for _, g := range Games {
		if g.ID == id {
			return g, nil
		}
	}
	return &Game{}, errors.New("Game with id of" + strconv.Itoa(id) + "not found")
}
