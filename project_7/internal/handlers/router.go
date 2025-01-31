package handlers

import (
	"learngo/httpgordle/internal/api"
	"learngo/httpgordle/internal/handlers/getstatus"
	"learngo/httpgordle/internal/handlers/guess"
	"learngo/httpgordle/internal/handlers/newgame"
	"learngo/httpgordle/internal/repository"
	"net/http"
)

// Mux creates a multiplexer with all the endpoints for our service.
// NewRouter returns a router that listens for requests to the following endpoints:
//   - Create a new game;
//
// The provided router is ready to serve.
func Mux(db *repository.GameRepository) *http.ServeMux {
	mux := http.NewServeMux()
	//  Connecting a URL to a handler
	// mux.HandleFunc(api.NewGameRoute, newgame.Handle)    // previous version
	mux.HandleFunc(http.MethodPost+" "+api.NewGameRoute, newgame.Handle(db))
	mux.HandleFunc(http.MethodGet+" "+api.GetStatusRoute, getstatus.Handle(db))
	mux.HandleFunc(http.MethodPut+" "+api.GuessRoute, guess.Handle(db))

	return mux
}
