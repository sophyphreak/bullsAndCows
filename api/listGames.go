package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../game"
	"github.com/gorilla/mux"
)

func listGames(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameID := vars["gameId"]
	gameIDInt, _ := strconv.Atoi(gameID)
	foundGames := game.ListGames()
	for _, game := range foundGames {
		if game.ID == gameIDInt {
			json.NewEncoder(w).Encode(game)
		}
	}
	json.NewEncoder(w).Encode(foundGames)
}
