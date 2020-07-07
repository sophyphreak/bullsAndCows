package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../game"
	"github.com/gorilla/mux"
)
<<<<<<< HEAD

type data struct {
	Rounds []game.Round `json:"rounds"`
}
=======
>>>>>>> 163ad410f8d1572b4331252c198a7abf1e195576

func listRounds(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	gameID := vars["gameId"]
<<<<<<< HEAD
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
=======
	//convert gameID to int
	gameIDInt, _ := strconv.Atoi(gameID)
	foundRounds := game.ListRounds(gameIDInt)
	json.NewEncoder(w).Encode(foundRounds)
>>>>>>> 163ad410f8d1572b4331252c198a7abf1e195576
}
