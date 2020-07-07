package api

import (
	"net/http"
<<<<<<< HEAD
)

func singleGame(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// gameId := vars["gameId"]
	// foundGame, err := game.GetGame(gameId)
	// check(err)
	// json.NewEncoder(w).Encode(foundGame)
=======

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
>>>>>>> 24f8b899984d104fc2079a3a7f50b5dff41116c8
}
