package gordle

import (
	"crypto/rand"
	_ "embed"
	"fmt"
	"math/big"
	"strings"
)

const (
	// ErrEmptyCorpus is returned when the provided corpus is empty.
	ErrEmptyCorpus = corpusError("corpus is empty")
	// ErrPickRandomWord is returned when a word has not been picked from the corpus.
	ErrPickRandomWord = corpusError("failed to pick a random word")
	ErrInaccessibleCorpus = corpusError("corpus can't be opened")
)

//go:embed corpus/english.txt
var corpus string

func ReadCorpus() ([]string, error) {
	// we expect the corpus to be a line- or space-separated list of words
	words := strings.Fields(corpus)

	if len(words) == 0 {
		return nil, ErrEmptyCorpus
	}
	return words, nil

}

// PickWord returns a random word from the corpus
func PickWord(corpus []string) (string, error) {
	index, err := rand.Int(rand.Reader, big.NewInt(int64(len(corpus))))
	if err != nil {
		return "", fmt.Errorf("failed to rand index (%s): %w", err, ErrPickRandomWord)
	}
	return corpus[index.Int64()], nil
}
