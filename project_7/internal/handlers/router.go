package handlers

import (
	"net/http"
	"learngo/httpgordle/internal/api"
	"learngo/httpgordle/internal/handlers/newgame"
)

// Mux creates a multiplexer with all the endpoints for our service.
func Mux() *http.ServeMux {
	mux := http.NewServeMux()
	//  Connecting a URL to a handler
	mux.HandleFunc(api.NewGameRoute, newgame.Handle)            
	return mux
}

