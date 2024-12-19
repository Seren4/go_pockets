package gordle

import (
	"slices"
	"strings"
	"testing"
	"errors"
)

func TestGameAsk(t *testing.T) {
	tt := map[string]struct {
		input string
		want  []rune
	}{
		"5 chars en": {
			input: "HELLO",
			want:  []rune("HELLO"),
		},
		"5 characters in arabic": {
			input: "مرحبا",
			want:  []rune("مرحبا"),
		},
		"5 characters in japanese": {
			input: "こんにちは",
			want:  []rune("こんにちは"),
		},
		"3 characters in japanese": {
			input: "こんに\nこんにちは",
			want:  []rune("こんにちは"),
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g := New(strings.NewReader(tc.input), string(tc.want), 0)
			got := g.ask()
			if !slices.Equal(got, tc.want) {
				t.Errorf("got = %v, want = %v", string(got), string(tc.want))
			}
		})
	}
}

func TestGameValidateGuest(t *testing.T) {
	tt := map[string]struct {
		word  []rune
		expected error
	}{
		"nominal": {
			word:  []rune("HELLO"),
			expected: nil,
		},
		"too long": {
			word:  []rune("BONJOUR"),
			expected: errInvalidWordlLength,
		},
		"too short": {
			word:  []rune("CIAO"),
			expected: errInvalidWordlLength,
		},
		"empty": {
			word:  []rune(""),
			expected: errInvalidWordlLength,
		},
		"nil": {
			word:  []rune(nil),
			expected: errInvalidWordlLength,
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g := New(nil, "HELLO", 0)
			err := g.validateGuess(tc.word)
			if !errors.Is(err, tc.expected){
				t.Errorf("%c got %q, expected %q", tc.word, tc.expected, err)
			}
		})
	}
}

func TestGameSplitToUpperCaseChars(t *testing.T) {
	tt := map[string]struct {
		word  string
		expected []rune
		} {
		"nominal": {
			word:  "hello",
			expected: []rune("HELLO"),
		},
		"upper and loxer": {
			word:  "helLO",
			expected: []rune("HELLO"),
		},
		"empty": {
			word:  "",
			expected: []rune(""),
		},
		"upper": {
			word:  "HELLO",
			expected: []rune("HELLO"),
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := splitToUpperCaseChars(string(tc.word))
			if !slices.Equal(got, tc.expected) {
				t.Errorf("got %v, expected %v", got, tc.expected)
			}
		})
	}
}
