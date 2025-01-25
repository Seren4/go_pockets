package main

import (
	"net/http"
	"fmt"
	"os"
)

func main() {
	// Start the server.
	// func ListenAndServe(addr string, handler Handler) error
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)                                             
	}
}