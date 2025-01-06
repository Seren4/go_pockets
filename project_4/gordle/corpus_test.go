package gordle_test

import (
	"learngo-pockets/gordle/gordle"
	"testing"
)

func TestReadCorpus(t *testing.T) {
	tt := map[string]struct {
		file string
		length  int
		err error
	}{
		"file ok": {
			file: "../corpus/english.txt",
			length: 14,
			err: nil,
		},
		"empty file": {
			file: "../corpus/empty.txt",
			length: 0,
			err: gordle.ErrCorpusIsEmpty,
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := gordle.ReadCorpus(tc.file)
			if tc.err != err {
				t.Errorf("got error %v, expected error %v", err, tc.err)

			}
			if len(got) != tc.length {
				t.Errorf("got %v, expected %v", len(got), tc.length)
			}
		})
	}
}