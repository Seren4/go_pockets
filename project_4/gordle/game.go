package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"slices"
)

// Game holds all the information we need to play a game of Gordle.
type Game struct{
	reader *bufio.Reader
	solution []rune
	maxAttempts int
}

// New returns a Game variable, which can be used to play.
func New(playerInput io.Reader, solution string, maxAttempts int) *Game {
	g := &Game{
		reader: bufio.NewReader(playerInput),
		solution: splitToUpperCaseChars(solution),
		maxAttempts: maxAttempts,
	}
	return g
}

//ask reads input until a valid suggestion is made (and returned).
func (g *Game) ask() []rune {
	fmt.Printf("Enter %d-character guess:\n", len(g.solution))
	for {
		playerInput, _, err := g.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Gordle failed to read your guess %s\n", err.Error())
			continue
		}
		guess := splitToUpperCaseChars(string(playerInput))

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
	if len(guess) != len(g.solution) {
		return fmt.Errorf("expected %d characters, got %d, %w", g.solution, len(guess), errInvalidWordlLength)

	} 
	return nil
}

// splitToUpperCaseChars turns a string into a list of characters
func splitToUpperCaseChars(input string) []rune {
	return []rune(strings.ToUpper(input))
}

// Play runs the game.
// ( creating a method on an object is achieved by writing a pointer receiver on the Game stucture)
func (g *Game) Play()  {
	fmt.Println("Welcome to Gordie")
	// ask for a valid word
	for i := 1; i <= g.maxAttempts; i++ {
		guess := g.ask()
		if slices.Equal(guess, g.solution) {
			fmt.Printf("You won! You find it in %d guess(es)\n", i)
			return
		}

	}
	fmt.Printf("🥲 You lost, the solution was: %s\n", string(g.solution))

}