package game

import (
	"errors"
	"strconv"
)

// Games contains all games, active and finished
var Games []*Game

// Game contains a single game's data
type Game struct {
	ID     int     `json:"id"`
	Answer [4]int  `json:"answer"`
	Status string  `json:"status"`
	Rounds []Round `json:"rounds"`
}

func findGameByID(id int) (int, error) {
	for i, g := range Games {
		if g.ID == id {
			return i, nil
		}
	}
	return -1, errors.New("Game with id of" + strconv.Itoa(id) + "not found")
}
