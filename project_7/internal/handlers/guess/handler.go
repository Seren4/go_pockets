package guess

import (
	"encoding/json"
	"learngo/httpgordle/internal/api"
	"learngo/httpgordle/internal/session"
	"learngo/httpgordle/internal/gordle"

	"log"
	"net/http"
	"fmt"
	"errors"
)

// The benefit of this method over http.Handle is that we don’t have to
// write a new http.Handler - we simply have to provide the handler itself,
// the function in charge of dealing with the request and writing the response.

// func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))


type gameGuesser interface {
	Find(session.GameID) (session.Game, error)
	Update(game session.Game) error
}

func Handle(db gameGuesser) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		id := req.PathValue(api.GameID)
		if id == "" {
			http.Error(w, "missing the ID of the game", http.StatusBadRequest)
			return
		}
		// TODO verify id
		// log package is not thread- safe.
		// So keep in mind that it can lead to unordered logs and complicate later testing.
		log.Printf("retrieve status of the game with id: %v", id)

		// Read the request, containing the guess, from the body of the input.
		// Instantiate the request struct
		r := api.GuessRequest{}
		// Decode the body
		err := json.NewDecoder(req.Body).Decode(&r)
		if err != nil {
			// If the JSON doesn’t parse, it’s a bad request
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		game, error := guess(session.GameID(id), r.Guess, db)
		if error != nil {
			switch {
				case errors.Is(err, session.ErrNotFound):
					http.Error(w, err.Error(), http.StatusNotFound)
				case errors.Is(err, gordle.ErrInvalidWordlLength):
					http.Error(w, err.Error(), http.StatusBadRequest)
				case errors.Is(err, session.ErrGameOver):
					http.Error(w, err.Error(), http.StatusForbidden)
				default:
					http.Error(w, err.Error(), http.StatusInternalServerError)
				}
				return
			}
	
		apiGame := api.ToGameResponse(game)
		w.Header().Set("Content-Type", "application/json")

		// Encode the game into JSON
		errorApiGame := json.NewEncoder(w).Encode(apiGame)
		if errorApiGame != nil {
			// The header has already been set. Nothing much we can do here.
			log.Printf("failed to write response: %s", err)
		}

	}
}

func guess(id session.GameID, guess string, db gameGuesser) (session.Game, error) {
	game, err := db.Find(id)  
	if game.AttemptsLeft == 0 || game.Status == session.StatusWon { 
		return session.Game{}, session.ErrGameOver
	}
	if err != nil {
		log.Printf("failed to find the game %s: %v", id, err)
		return session.Game{}, fmt.Errorf("can't find the game: %w", err)
	}

	fb, err := game.Gordle.Play(guess)

	if err != nil {
		return session.Game{}, fmt.Errorf("unable to play: %w", err)
	}

	game.Guesses = append(game.Guesses, session.Guess{
		Word:     guess,
		Feedback: fb.String(),
	})

	game.AttemptsLeft -= 1

	switch {                                                      
		case fb.GameWon():
			game.Status = session.StatusWon
		case game.AttemptsLeft == 0:
			game.Status = session.StatusLost
		default:
			game.Status = session.StatusPlaying
		}

	err = db.Update(game) 
	if err != nil {
		return session.Game{}, fmt.Errorf("unable to save game: %w", err)
	}
	return game, nil

}
