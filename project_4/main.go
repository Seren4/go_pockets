package main

import (
	"fmt"
	"learngo-pockets/gordle/gordle"
	"os"
)

const maxAttempts = 3

func main()  {
	
	corpus, err := gordle.ReadCorpus("corpus/english.txt")
	if err != nil {
		_,_ = fmt.Fprintf(os.Stderr, "unable to read corpus file: %s", err)
		return
	}
	
	g, err := gordle.New(os.Stdin, corpus, maxAttempts)
	if err != nil {
		_,_ = fmt.Fprintf(os.Stderr, "unable to start the game: %s", err)
		return

	}
	g.Play()
}