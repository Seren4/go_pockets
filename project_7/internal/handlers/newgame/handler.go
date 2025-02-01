package newgame

import (
	"encoding/json"
	"fmt"
	"learngo/httpgordle/internal/api"
	"learngo/httpgordle/internal/gordle"
	"learngo/httpgordle/internal/session"
	"log"
	"net/http"

	"github.com/oklog/ulid/v2"
)

const maxAttempts = 3

// The benefit of this method over http.Handle is that we don’t have to
// write a new http.Handler - we simply have to provide the handler itself,
// the function in charge of dealing with the request and writing the response.

// func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))

// The NewGame endpoint only needs to add a game to the repository, nothing else. We can
// actually prevent it from doing anything else by defining a minimal interface.
type gameAdder interface {
	Add(game session.Game) error
}

// Let’s anonymise the Handle functions and wrap them instead in a Handler function that
// takes a repository as a parameter and returns the previous http.HandleFunc.
func Handle(db gameAdder) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		game, err := createGame(db)
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
}

func createGame(db gameAdder) (session.Game, error) {
	corpus, err := gordle.ReadCorpus()
	if err != nil {
		return session.Game{}, fmt.Errorf("unable to read corpus: %w", err)
	}
	game, err := gordle.New(corpus)
	if err != nil {
		return session.Game{}, fmt.Errorf("unable to create a new game: %w", err)
	}
	g := session.Game{
		ID:           session.GameID(ulid.Make().String()),
		Gordle:       *game,
		Guesses:      []session.Guess{},
		Status:       session.StatusPlaying,
		AttemptsLeft: maxAttempts,
	}
	err = db.Add(g)
	if err != nil {
		return session.Game{}, fmt.Errorf("failed to save the new game")
	}
	return g, nil
}
