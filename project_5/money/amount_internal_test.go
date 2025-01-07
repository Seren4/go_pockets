package money

import (
	"errors"
	"testing"
)

func TestNewAmount(t *testing.T) {
	tt := map[string]struct {
		quantity  Decimal
		currency Currency
		expected byte
		err      error
	}{
		"precision ok 1": {
			quantity:  Decimal{subunits: 34056, precision: 1},
			currency:  Currency{code: "IRR", precision: 2},
			expected: 2,
			err:      nil,
		},
		"precision ok from 0 to 2": {
			quantity:  Decimal{subunits: 34056, precision: 0},
			currency:  Currency{code: "IRR", precision: 2},
			expected: 2,
			err:      nil,
		},
		"precision ok from 2 to 2": {
			quantity:  Decimal{subunits: 34056, precision: 2},
			currency:  Currency{code: "IRR", precision: 2},
			expected: 2,
			err:      nil,
		},
		"precision too big": {
			quantity:  Decimal{subunits: 34056, precision: 3},
			currency:  Currency{code: "IRR", precision: 2},
			expected: 0,
			err:      ErrDecimalTooPrecise,
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			amount, err := NewAmount(tc.quantity, tc.currency)
			if amount.quantity.precision != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, amount.quantity.precision)
			}
			if !errors.Is(err, tc.err) {
				t.Errorf("expected error %v, got %v", tc.err, err)
			}
		})
	}
}

