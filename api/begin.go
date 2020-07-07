package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../game"
)

type b struct {
	ID int `json:"id"`
}

func begin(w http.ResponseWriter, r *http.Request) {
	id := game.Begin()
	fmt.Println("POST /begin")
	json.NewEncoder(w).Encode(b{id})
}
