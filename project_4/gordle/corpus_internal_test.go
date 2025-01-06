package gordle_test

import (
	"learngo-pockets/gordle/gordle"
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
		file string
		word  string
		expected bool
	}{
		"word in corpus": {
			file: "../corpus/english.txt",
			word: "SALIENT",
			expected: true,
		},
		"empty file": {
			file: "../corpus/empty.txt",
			word: "ERRORS",
			expected: false,
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, _ := gordle.ReadCorpus(tc.file)
			result := inCorpus(got, tc.word)
			if tc.expected != result {
				t.Errorf("got error")
			}
		})
	}
}