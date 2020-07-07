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
	fmt.Println("POST /begin")
	id := game.Begin()
	json.NewEncoder(w).Encode(b{id})
}
