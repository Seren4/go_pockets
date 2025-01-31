package gordle_test

import (
	"testing"
)

func inCorpus(corpus []string, word string) bool {
	for _, corpusWord := range corpus {
		if corpusWord == word {
			return true
		}
	}
	return false
}

func TestPickWord(t *testing.T) {
	tt := map[string]struct {
		words    []string
		word     string
		expected bool
	}{
		"word in corpus": {
			word:     "ПРИВЕТ",
			words:    []string{"HELLO", "SALUT", "ПРИВЕТ", "ΧΑΙΡΕ"},
			expected: true,
		},
		"empty file": {
			words:    []string{"HELLO", "SALUT", "ПРИВЕТ", "ΧΑΙΡΕ"},
			word:     "ERRORS",
			expected: false,
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			result := inCorpus(tc.words, tc.word)
			if tc.expected != result {
				t.Errorf("got error")
			}
		})
	}
}
