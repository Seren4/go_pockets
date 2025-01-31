package gordle

import (
	"fmt"
	"os"
	"strings"
)

// Game holds all the information we need to get the Feedback of a play.
type Game struct {
	solution []rune
}

// New returns a Game variable, which can be used to play.
func New(solution string) (*Game, error) {
	if len(solution) == 0 {
		return nil, ErrEmptyCorpus
	}
	g := &Game{
		solution: splitToUpperCaseChars(solution),
	}
	return g, nil
}

// errInvalidWordlLength
var errInvalidWordlLength = gameError("invalid guess length, word doesn't have the same nb of chars as the solution")

// validateGuess ensures the guess is valid enough.
func (g *Game) validateGuess(guess string) error {
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
func (g *Game) Play(guess string) (Feedback, error) {
	err := g.validateGuess(guess)
	if err != nil {
		return Feedback{}, fmt.Errorf("this guess is not the correct length: %w", err)
	}
	characters := splitToUpperCaseChars(guess)
	fb := computeFeedback(characters, g.solution)

	return fb, nil

}

// ShowAnswer gives up on playing this game. It returns the solution.
func (g *Game) ShowAnswer() string {
	return string(g.solution)
}

// computeFeedback verifies every character of the guess against the solution
func computeFeedback(guess, solution []rune) Feedback {
	// initialize holders for marks
	fb := make(Feedback, len(solution))
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
