package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Game holds all the information we need to play a game of Gordle.
type Game struct{
	reader *bufio.Reader
}

// New returns a Game variable, which can be used to play.
func New(playerInput io.Reader) *Game {
	g := &Game{
		reader: bufio.NewReader(playerInput),
	}
	return g
}

const solutionLength = 5

//ask reads input until a valid suggestion is made (and returned).
func (g *Game) ask() []rune {
	fmt.Printf("Enter %d-character guess:\n", solutionLength)
	for {
		playerInput, _, err := g.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Gordle failed to read your guess %s\n", err.Error())
			continue
		}
		guess := []rune(string(playerInput))

		if len(guess) != solutionLength {
			_, _ = fmt.Fprintf(os.Stderr, "Invali attempt: expected %d characters, got %d\n", solutionLength, len(guess))

		} else {
			return guess
		}
	}
}

// Play runs the game.
// ( creating a method on an object is achieved by writing a pointer receiver on the Game stucture)
func (g *Game) Play()  {
	fmt.Println("Welcome to Gordie")
	// ask for a valid word
	guess := g.ask()
	fmt.Printf("Your guess is %s\n", string(guess))

}