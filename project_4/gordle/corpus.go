package gordle

import (
	"fmt"
	"os"
	"strings"
)

func ReadCorpus(path string) ([]string, error) {

	corpuslist, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Unable to open %q: %w", path, err)
	}

	if len(corpuslist) == 0 {
		return nil, fmt.Errorf("Error, corpus %q is empty", path)
	}
	words := strings.Fields(string(corpuslist))
	return words, nil

}
