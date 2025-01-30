package newgame

import (
	"encoding/json"
	"learngo/httpgordle/internal/api"
	"learngo/httpgordle/internal/session"
	"log"
	"net/http"
)

// The benefit of this method over http.Handle is that we don’t have to
// write a new http.Handler - we simply have to provide the handler itself,
// the function in charge of dealing with the request and writing the response.

// func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))

func Handle(w http.ResponseWriter, req *http.Request) {

	game, err := createGame()
	if err != nil {
		log.Printf("unable to create a new game: %s", err)
		http.Error(w, "failed to create a new game", http.StatusInternalServerError)
		return
	}

	// Tell the consumer that we are sending JSON
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	apiGame := api.ToGameResponse(game)
	// Transform the game into the API format
	// Encode the game into JSON
	error := json.NewEncoder(w).Encode(apiGame)
	if error != nil {
		// The header has already been set. Nothing much we can do here.
		log.Printf("failed to write response: %s", err)
	}
}

func createGame() (session.Game, error) {
	return session.Game{}, nil
}
