package gordle

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

// Game holds all the information we need to play a game of Gordle.
type Game struct {
	reader      *bufio.Reader
	solution    []rune
	maxAttempts int
}

// New returns a Game variable, which can be used to play.
func New(playerInput io.Reader, corpus []string, maxAttempts int) (*Game, error) {
	if len(corpus) == 0 {
		return nil, ErrCorpusIsEmpty
	}
	g := &Game{
		reader:      bufio.NewReader(playerInput),
		solution:    splitToUpperCaseChars(PickWord(corpus)),
		maxAttempts: maxAttempts,
	}
	return g, nil
}

// ask reads input until a valid suggestion is made (and returned).
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
		return fmt.Errorf("expected %d characters, got %d, %w", len(g.solution), len(guess), errInvalidWordlLength)

	}
	return nil
}

// splitToUpperCaseChars turns a string into a list of characters
func splitToUpperCaseChars(input string) []rune {
	return []rune(strings.ToUpper(input))
}

// Play runs the game.
// ( creating a method on an object is achieved by writing a pointer receiver on the Game stucture)
func (g *Game) Play() {
	fmt.Println("Welcome to Gordie")
	// ask for a valid word
	for i := 1; i <= g.maxAttempts; i++ {
		guess := g.ask()
		fb := computeFeedback(guess, g.solution)
		fmt.Println(fb)

		if slices.Equal(guess, g.solution) {
			fmt.Printf("You won! You find it in %d guess(es)\n", i)
			return
		}

	}
	fmt.Printf("🥲 You lost, the solution was: %s\n", string(g.solution))

}

// computeFeedback verifies every character of the guess against the solution 
func computeFeedback(guess, solution []rune) feedback {
	// initialize holders for marks
	fb := make(feedback, len(solution))
	// checked keeps trace of the already checked character in the solution
	checked := make([]bool, len(solution))

	// repeate the length check just in case
	if len(guess) != len(solution) {
		_, _ = fmt.Fprintf(os.Stderr, "Internal error, guess and solution have differents lengths: %d vs %d\n", len(guess), len(solution))
		return fb
	}

	// loop n. 1: check for chars in correct position
	for index, letter := range guess {
		if letter == solution[index] {
			fb[index] = correctPosition
			checked[index] = true
		}
	}
	
	// loop n. 2: check for chars in wrong position, only if the not already checked
	for index, letter := range guess {
		if fb[index] != absentChar {
			// the char has already been marked.
			continue
		}

		for indexSol, letterSol := range solution {
			// already checked, go on
			if checked[indexSol] {
				continue
			}
			if letterSol == letter {
				fb[index] = wrongPosition
				checked[indexSol] = true
				// Skip to the next letter
				break
				}
			}

		}
	return fb
}
