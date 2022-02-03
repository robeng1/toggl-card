package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/robeng1/toggl/api"
	"github.com/robeng1/toggl/storage/memory"
)

func serve() {
	store := memory.NewStore()
	handler := api.NewHandler(store)
	// creates a new instance of a mux router
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/decks", handler.CreateDeck).Methods("POST")
	router.HandleFunc("/decks/{id}", handler.OpenDeck).Methods("GET")
	router.HandleFunc("/decks/{id}/draw/{count}", handler.DrawCard).Methods("GET")
	router.HandleFunc("/", handler.HealthCheck).Methods("GET")
	log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
	serve()
}
