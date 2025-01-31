package main

import (
	"fmt"
	"learngo/httpgordle/internal/handlers"
	"learngo/httpgordle/internal/repository"
	"net/http"
	"os"
)


func main() {
	// Init the data storage.
	db := repository.New()
	// Start the server.
	// func ListenAndServe(addr string, handler Handler) error
	err := http.ListenAndServe(":8080", handlers.Mux(db))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)                                             
	}
}
