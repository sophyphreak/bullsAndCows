package api

import (
	"encoding/json"
	"net/http"

	"../game"
)

type b struct {
	ID int `json:"id"`
}

func begin(w http.ResponseWriter, r *http.Request) {
	id := game.Begin()
	json.NewEncoder(w).Encode(b{id})
}
