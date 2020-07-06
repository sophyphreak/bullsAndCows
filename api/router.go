package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/begin", begin).Methods("POST")
	router.HandleFunc("/guess", makeGuess).Methods("POST")
	router.HandleFunc("/game", listGames).Methods("GET")
	router.HandleFunc("/game/{gameId}", singleGame).Methods("GET")
	router.HandleFunc("/rounds/{gameId}", listRounds).Methods("GET")
	log.Fatal(http.ListenAndServe(":10000", router))
}
