package gordle

import (
	"testing"
)

func TestGameFeedbackString(t *testing.T) {
	const newHint hint = 4
	tt := map[string]struct {
		fb   feedback
		want string
	}{
		"one absent": {
			fb:   feedback{correctPosition, correctPosition, correctPosition, correctPosition, absentChar},
			want: "✅✅✅✅Ⅹ",
		},
		"invalid hint": {
			fb:   feedback{correctPosition, correctPosition, correctPosition, correctPosition, newHint},
			want: "✅✅✅✅⚠️",
		},
		"all symbols": {
			fb:   feedback{correctPosition, wrongPosition, absentChar},
			want: "✅🟡Ⅹ",
		},
		"only absent": {
			fb:   feedback{absentChar, absentChar, absentChar, absentChar, absentChar},
			want: "ⅩⅩⅩⅩⅩ",
		},
		"empty list": {
			fb:   feedback{},
			want: "",
		},
		"nil list": {
			fb:   nil,
			want: "",
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := tc.fb.String()
			if got != tc.want {
				t.Errorf("got %v, expected %v", got, tc.want)
			}
		})
	}
}
