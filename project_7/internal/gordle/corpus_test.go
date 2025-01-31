package gordle_test

import (
	"learngo/httpgordle/internal/gordle"
	"testing"
)

func TestParseCorpus(t *testing.T) {
	words, err := gordle.ReadCorpus()
	if err != nil {
		t.Fatalf("expected no error, got %s", err)
	}

	const wordsInEnglishCorpus = 14
	if len(words) != wordsInEnglishCorpus {
		t.Errorf("expected %d words, got %d", wordsInEnglishCorpus, len(words))
	}
}
