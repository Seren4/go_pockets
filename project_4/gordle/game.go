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

		err = g.validateGuess(guess)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid: %s.\n", err.Error())
		} else {
			return guess
		}
	}
}

// errInvalidWordlLength
var errInvalidWordlLength = fmt.Errorf("invalid guess, word doesn't have the same nb of chars as the solution")

// validateGuess ensures the guess is valid enough.
func (g *Game) validateGuess(guess []rune) error {
	if len(guess) != solutionLength {
		return fmt.Errorf("expected %d characters, got %d, %w", solutionLength, len(guess), errInvalidWordlLength)

	} 
	return nil
}


// Play runs the game.
// ( creating a method on an object is achieved by writing a pointer receiver on the Game stucture)
func (g *Game) Play()  {
	fmt.Println("Welcome to Gordie")
	// ask for a valid word
	guess := g.ask()
	fmt.Printf("Your guess is %s\n", string(guess))

}