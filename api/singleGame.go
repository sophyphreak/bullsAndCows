package api

import (
	"encoding/json"
	"net/http"

	"../game"
	"github.com/gorilla/mux"
)

func singleGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameId := vars["gameId"]
	foundGame, err := game.GetGame(gameId)
	check(err)
	json.NewEncoder(w).Encode(foundGame)
}
