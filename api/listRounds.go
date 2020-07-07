package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../game"
	"github.com/gorilla/mux"
)

type data struct {
	Rounds []game.Round `json:"rounds"`
}

func listRounds(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameID := vars["gameId"]
	id, err := strconv.Atoi(gameID)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	g, err := game.GetGame(id)
	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}
	json.NewEncoder(w).Encode(g.Rounds)
}
