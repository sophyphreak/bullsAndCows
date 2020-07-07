package api

import (
	"encoding/json"
	"net/http"

	"strconv"

	"../game"
	"github.com/gorilla/mux"
)

func singleGame(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameId := vars["gameId"]
	//convert gameID to int
	gameIdInt, _ := strconv.Atoi(gameId)
	foundGame, err := game.GetGame(gameIdInt)
	game.Check(err)
	json.NewEncoder(w).Encode(foundGame)
}
