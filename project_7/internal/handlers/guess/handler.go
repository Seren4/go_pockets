package guess

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
		game, error := guess(id, r)
		if error != nil {
			// The header has already been set. Nothing much we can do here.
			log.Printf("failed to write response: %s", err)
		}
		apiGame := api.ToGameResponse(game)
		// Encode the game into JSON
		errorApiGame := json.NewEncoder(w).Encode(apiGame)
		if errorApiGame != nil {
			// The header has already been set. Nothing much we can do here.
			log.Printf("failed to write response: %s", err)
		}

	}
}

func guess(id string, r api.GuessRequest) (session.Game, error) {
	return session.Game{
		ID: session.GameID(id),
	}, nil
}
