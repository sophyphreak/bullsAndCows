package api

import (
	"encoding/json"
	"net/http"

	"../game"
)

func listGames(w http.ResponseWriter, r *http.Request) {
	foundGames := game.ListGames()
	json.NewEncoder(w).Encode(foundGames)
}
