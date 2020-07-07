package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"../game"
)

type guess struct {
	ID    int    `json:"id"`
	Guess string `json:"guess"`
}

func makeGuess(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var g guess
	json.Unmarshal(reqBody, &g)
	round, err := game.MakeGuess(g.ID, g.Guess)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(round)
}
