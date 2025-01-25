package main

import (
	"net/http"
	"fmt"
)

func main() {
	// Start the server.
	// func ListenAndServe(addr string, handler Handler) error
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		// It’s OK to panic in the body of main.
		panic(err)                                                   
	}
	fmt.Println(err)
}