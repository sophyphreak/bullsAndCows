package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"../game"
)

type guess struct {
	id    string `json:"id"`
	guess string `json:"guess"`
}

func makeGuess(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var g guess
	json.Unmarshal(reqBody, &g)
	round := game.MakeGuess(g.id, g.guess)
	json.NewEncoder(w).Encode(round)
}
