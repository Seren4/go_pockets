package money

import (
	"testing"
)

func TestApplyExchangeRate(t *testing.T) {
	tt := map[string]struct {
		amount  Amount
		target Currency
		rate ExchangeRate
		expected Amount
	}{
		"rate 1 same precision": {
			amount: Amount{quantity: Decimal{subunits: 2345, precision: 2}, currency: Currency{code: "EUR", precision: 2}},
			target: Currency{code: "IRR", precision: 2},
			rate: ExchangeRate{subunits: 1, precision: 0},
			expected: Amount{quantity: Decimal{subunits: 2345, precision: 2}, currency: Currency{code: "IRR", precision: 2}},
		},
		"rate 2 precision target > amout": {
			amount: Amount{quantity: Decimal{subunits: 2345, precision: 2}, currency: Currency{code: "EUR", precision: 2}},
			target: Currency{code: "IRR", precision: 3},
			rate: ExchangeRate{subunits: 2, precision: 0},
			expected: Amount{quantity: Decimal{subunits: 46900, precision: 3}, currency: Currency{code: "IRR", precision: 3}},
		},
		"rate 2 precision target < amout": {
			amount: Amount{quantity: Decimal{subunits: 2345, precision: 2}, currency: Currency{code: "EUR", precision: 2}},
			target: Currency{code: "IRR", precision: 1},
			rate: ExchangeRate{subunits: 2, precision: 0},
			expected: Amount{quantity: Decimal{subunits: 469, precision: 1}, currency: Currency{code: "IRR", precision: 1}},
		},
		"rate 4": {
			amount: Amount{quantity: Decimal{subunits: 250, precision: 2}, currency: Currency{code: "EUR", precision: 2}},
			target: Currency{code: "IRR", precision: 2},
			rate: ExchangeRate{subunits: 4, precision: 0},
			expected: Amount{quantity: Decimal{subunits: 1000, precision: 2}, currency: Currency{code: "IRR", precision: 2}},
		},
		"rate 2.5": {
			amount: Amount{quantity: Decimal{subunits: 4, precision: 0}, currency: Currency{code: "EUR", precision: 2}},
			target: Currency{code: "IRR", precision: 0},
			rate: ExchangeRate{subunits: 25, precision: 1},
			expected: Amount{quantity: Decimal{subunits: 10, precision: 0}, currency: Currency{code: "IRR", precision: 0}},
		},
		"rate 2.52678": {
			amount: Amount{quantity: Decimal{subunits: 314, precision: 2}, currency: Currency{code: "EUR", precision: 2}},
			target: Currency{code: "IRR", precision: 2},
			rate: ExchangeRate{subunits: 252678, precision: 5},
			expected: Amount{quantity: Decimal{subunits: 793, precision: 2}, currency: Currency{code: "IRR", precision: 2}},
		},
		"rate 10": {
			amount: Amount{quantity: Decimal{subunits: 11, precision: 1}, currency: Currency{code: "EUR", precision: 1}},
			target: Currency{code: "IRR", precision: 1},
			rate: ExchangeRate{subunits: 10, precision: 0},
			expected: Amount{quantity: Decimal{subunits: 110, precision: 1}, currency: Currency{code: "IRR", precision: 1}},
		},
		"large nb": {
			amount: Amount{quantity: Decimal{subunits: 100000000001, precision: 2}, currency: Currency{code: "EUR", precision: 2}},
			target: Currency{code: "IRR", precision: 2},
			rate: ExchangeRate{subunits: 2, precision: 0},
			expected: Amount{quantity: Decimal{subunits: 200000000002, precision: 2}, currency: Currency{code: "IRR", precision: 2}},
		},
		"large nb 2": {
			amount: Amount{quantity: Decimal{subunits: 265413, precision: 0}, currency: Currency{code: "EUR", precision: 0}},
			target: Currency{code: "IRR", precision: 3},
			rate: ExchangeRate{subunits: 1, precision: 0},
			expected: Amount{quantity: Decimal{subunits: 265413000, precision: 3}, currency: Currency{code: "IRR", precision: 3}},
		},
		"precision": {
			amount: Amount{quantity: Decimal{subunits: 2, precision: 0}, currency: Currency{code: "EUR", precision: 0}},
			target: Currency{code: "IRR", precision: 5},
			rate: ExchangeRate{subunits: 1337, precision: 3},
			expected: Amount{quantity: Decimal{subunits: 267400, precision: 5}, currency: Currency{code: "IRR", precision: 5}},
		},
		"precision2": {
			amount: Amount{quantity: Decimal{subunits: 2, precision: 0}, currency: Currency{code: "EUR", precision: 0}},
			target: Currency{code: "IRR", precision: 5},
			rate: ExchangeRate{subunits: 133, precision: 18},
			expected: Amount{quantity: Decimal{subunits: 0, precision: 5}, currency: Currency{code: "IRR", precision: 5}},
		},
		
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got := applyExchangeRate(tc.amount, tc.target, tc.rate)
			if got != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}
