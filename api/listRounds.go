package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../game"
	"github.com/gorilla/mux"
)

func listRounds(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameID := vars["gameId"]
	//convert gameID to int
	gameIDInt, _ := strconv.Atoi(gameID)
	foundRounds := game.ListRounds(gameIDInt)
	json.NewEncoder(w).Encode(foundRounds)
}
