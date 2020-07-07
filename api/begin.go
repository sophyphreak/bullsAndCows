package api

import (
	"encoding/json"
	"net/http"
)

type b struct {
	ID int `json:"id"`
}

func begin(w http.ResponseWriter) {
	id := game.Begin()
	json.NewEncoder(w).Encode(b{id})
}
