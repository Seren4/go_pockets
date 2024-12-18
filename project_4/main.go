package main

import (
	"os"
	"learngo-pockets/gordle/gordle"
)

const maxAttempts = 3

func main()  {
	solution := "pocket"
	g := gordle.New(os.Stdin, solution, maxAttempts)
	g.Play()
}