package gordle

import (
	"errors"
	"slices"
	"testing"
)

func TestGameValidateGuest(t *testing.T) {
	tt := map[string]struct {
		word     string
		expected error
	}{
		"nominal": {
			word:     "HELLO",
			expected: nil,
		},
		"too long": {
			word:     "BONJOUR",
			expected: ErrInvalidWordlLength,
		},
		"too short": {
			word:     "CIAO",
			expected: ErrInvalidWordlLength,
		},
		"empty": {
			word:     "",
			expected: ErrInvalidWordlLength,
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			g, _ := New("HELLO")
			err := g.validateGuess(tc.word)
			if !errors.Is(err, tc.expected) {
				t.Errorf("%s got %q, expected %q", tc.word, tc.expected, err)
			}
		})
	}
}

func TestGameSplitToUpperCaseChars(t *testing.T) {
	tt := map[string]struct {
		word     string
		expected []rune
	}{
		"nominal": {
			word:     "hello",
			expected: []rune("HELLO"),
		},
		"upper and loxer": {
			word:     "helLO",
			expected: []rune("HELLO"),
		},
		"empty": {
			word:     "",
			expected: []rune(""),
		},
		"upper": {
			word:     "HELLO",
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

func TestGamecomputeFeedback(t *testing.T) {
	tt := map[string]struct {
		guess            string
		solution         string
		expectedFeedback Feedback
	}{
		"nominal": {
			guess:            "hello",
			solution:         "hello",
			expectedFeedback: Feedback{correctPosition, correctPosition, correctPosition, correctPosition, correctPosition},
		},
		"one char": {
			guess:            "qwrtl",
			solution:         "hello",
			expectedFeedback: Feedback{absentChar, absentChar, absentChar, absentChar, wrongPosition},
		},
		"different length": {
			guess:            "hello",
			solution:         "bonjour",
			expectedFeedback: Feedback{absentChar, absentChar, absentChar, absentChar, absentChar, absentChar, absentChar},
		},
		"double character with wrong answer": {
			guess:            "helll",
			solution:         "hello",
			expectedFeedback: Feedback{correctPosition, correctPosition, correctPosition, correctPosition, absentChar},
		},
		"five identical, two ok": {
			guess:            "lllll",
			solution:         "hello",
			expectedFeedback: Feedback{absentChar, absentChar, correctPosition, correctPosition, absentChar},
		},
		"two identical, different position": {
			guess:            "hlleo",
			solution:         "hello",
			expectedFeedback: Feedback{correctPosition, wrongPosition, correctPosition, wrongPosition, correctPosition},
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := computeFeedback([]rune(tc.guess), []rune(tc.solution))
			// if !slices.Equal(got, tc.expectedFeedback) {
			// 	t.Errorf("got %v, expected %v", got, tc.expectedFeedback)
			// }
			// OR with Equal
			if !got.Equal(tc.expectedFeedback) {
				t.Errorf("got %v, expected %v", got, tc.expectedFeedback)

			}
		})
	}
}
