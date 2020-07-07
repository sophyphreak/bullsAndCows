package api

import (
	"encoding/json"
	"net/http"

	"../game"
	"github.com/gorilla/mux"
)

func listGames(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["game"]
	foundGames := game.ListGames()
	for _, game := range foundGames {
		if game.ID == key {
			json.NewEncoder(w).Encode(account)
		}
	}
	foundGames := game.ListGames()
	json.NewEncoder(w).Encode(foundGames)
}
