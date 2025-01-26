package main

import (
	"net/http"
	"fmt"
	"os"
	"learngo/httpgordle/internal/handlers"
)


func main() {
	// Start the server.
	// func ListenAndServe(addr string, handler Handler) error
	err := http.ListenAndServe(":8080", handlers.Mux())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)                                             
	}
}
