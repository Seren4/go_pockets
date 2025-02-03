package getstatus

import (
	"encoding/json"
	"learngo/httpgordle/internal/api"
	"learngo/httpgordle/internal/session"
	"log"
	"net/http"
	"errors"
)

// The benefit of this method over http.Handle is that we don’t have to
// write a new http.Handler - we simply have to provide the handler itself,
// the function in charge of dealing with the request and writing the response.

// func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
type gameFinder interface {
	Find(gameID session.GameID) (session.Game, error)
}

func Handle(db gameFinder) http.HandlerFunc {
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
		
		game, err := db.Find(session.GameID(id))            
		if err != nil {
			if errors.Is(err, session.ErrNotFound) {
				http.Error(w, "this game does not exist", http.StatusNotFound)
				return
			}
			log.Printf("cannot fetch game %s: %s", id, err)
			http.Error(w, "failed to fetch game", http.StatusInternalServerError)
			return
		}

		apiGame := api.ToGameResponse(game)
		// Encode the game into JSON
		error := json.NewEncoder(w).Encode(apiGame)
		if error != nil {
			// The header has already been set. Nothing much we can do here.
			log.Printf("failed to write response: %s", err)
		}

	}
}
