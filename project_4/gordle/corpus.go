package gordle

import (
	"fmt"
	"os"
	"strings"
	"golang.org/x/exp/rand"
)

const ErrCorpusIsEmpty = corpusError("corpus is empty")

func ReadCorpus(path string) ([]string, error) {

	corpuslist, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Unable to open %q: %w", path, err)
	}

	if len(corpuslist) == 0 {
		return nil, ErrCorpusIsEmpty
	}
	words := strings.Fields(string(corpuslist))
	return words, nil

}

// PickWord returns a random word from the corpus 
func PickWord(corpus []string) string{
	index := rand.Intn(len(corpus))
	return corpus[index]
}