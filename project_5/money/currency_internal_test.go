package money

import (
	"errors"
	"testing"
)

func TestParseCurrency(t *testing.T) {
	tt := map[string]struct {
		code  string
		expected Currency
		err      error
	}{
		"code ok": {
			code:  "IRR",
			expected: Currency{code: "IRR", precision: 0},
			err:      nil,
		},
		"code too long": {
			code:  "IRRQ",
			expected: Currency{code: "", precision: 0},
			err:      ErrInvalidCurrencyCode,
		},
		"code too short": {
			code:  "IR",
			expected: Currency{code: "", precision: 0},
			err:      ErrInvalidCurrencyCode,
		},
		"code with symbols": {
			code:  "IR%",
			expected: Currency{code: "", precision: 0},
			err:      ErrInvalidCurrencyCode,
		},
		"no code": {
			code:  "",
			expected: Currency{code: "", precision: 0},
			err:      ErrInvalidCurrencyCode,
		},
		"code not capitalized": {
			code:  "aaa",
			expected: Currency{code: "aaa", precision: 2},
			err:      nil,
		},
		
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			gotCurrency, gotError := ParseCurrency(tc.code)
			if gotCurrency != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, gotCurrency)
			}
			if !errors.Is(gotError, tc.err) {
				t.Errorf("expected error %v, got %v", tc.err, gotError)

			}
		})
	}
}
