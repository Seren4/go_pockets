package money

import (
	"errors"
	"testing"
)

func TestParseDecimal(t *testing.T) {
	tt := map[string]struct {
		decimal  string
		expected Decimal
		err      error
	}{
		"two digits": {
			decimal:  "34.98",
			expected: Decimal{subunits: 3498, precision: 2},
			err:      nil,
		},
		"three digits": {
			decimal:  "34.989",
			expected: Decimal{subunits: 34989, precision: 3},
			err:      nil,
		},
		"NO digits": {
			decimal:  "",
			expected: Decimal{subunits: 0, precision: 0},
			err:      ErrInvalidDecimal,
		},
		"Too big": {
			decimal:  "1234567890123",
			expected: Decimal{subunits: 0, precision: 0},
			err:      ErrTooLarge,
		},
		"with 0 as after ": {
			decimal:  "10.0",
			expected: Decimal{subunits: 10, precision: 0},
			err:      nil,
		},
		"with 0 as before ": {
			decimal:  "0.556",
			expected: Decimal{subunits: 556, precision: 3},
			err:      nil,
		},
		"int": {
			decimal:  "22",
			expected: Decimal{subunits: 22, precision: 0},
			err:      nil,
		},
		"zero": {
			decimal:  "0",
			expected: Decimal{subunits: 0, precision: 0},
			err:      nil,
		},
		"comma as separator": {
			decimal:  "34,989",
			expected: Decimal{subunits: 0, precision: 0},
			err:      ErrInvalidDecimal,
		},
		"random symbol as separator": {
			decimal:  "34-989",
			expected: Decimal{subunits: 0, precision: 0},
			err:      ErrInvalidDecimal,
		},
		"NaN": {
			decimal:  "test",
			expected: Decimal{subunits: 0, precision: 0},
			err:      ErrInvalidDecimal,
		},
		"invalid decimal part": {
			decimal:  "20.money",
			expected: Decimal{subunits: 0, precision: 0},
			err:      ErrInvalidDecimal,
		},
		"invalid unit part": {
			decimal:  "money.20",
			expected: Decimal{subunits: 0, precision: 0},
			err:      ErrInvalidDecimal,
		},
		"prefix 0 as decimal digits": {
			decimal:  "34.056",
			expected: Decimal{subunits: 34056, precision: 3},
			err: nil,
		},
		"sufix 0 as decimal digits": {
			decimal:  "34.50",
			expected: Decimal{subunits: 345, precision: 1},
			err: nil,
		},
		
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			gotDecimal, gotError := ParseDecimal(tc.decimal)
			if gotDecimal != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, gotDecimal)
			}
			if !errors.Is(gotError, tc.err) {
				t.Errorf("expected error %v, got %v", tc.err, gotError)

			}
		})
	}
}


func TestStringDecimal(t *testing.T) {
	tt := map[string]struct {
		expected  string
		decimal Decimal
	}{
		"precision 0": {
			decimal:  Decimal{precision: 0, subunits: 555},
			expected: "555",
		},
		"precision 1": {
			decimal:  Decimal{precision: 1, subunits: 55},
			expected: "5.5",
		},
		"precision 2": {
			decimal:  Decimal{precision: 2, subunits: 555},
			expected: "5.55",
		},
		"precision 3": {
			decimal:  Decimal{precision: 3, subunits: 505000},
			expected: "505.000",
		},
		"precision 4": {
			decimal:  Decimal{precision: 4, subunits: 55505},
			expected: "5.5505",
		},
		"precision 4 bis": {
			decimal:  Decimal{precision: 4, subunits: 505},
			expected: "0.0505",
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := tc.decimal.String()
			if got != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		
		})
	}
}