package ecbank

import (
	"testing"
	"learngo-pockets/moneyconverter/money"
)

func TestExchangeRate(t *testing.T) {
	tt := map[string]struct {
		envelope envelope
		source  string
		target string
		err error
		expected money.ExchangeRate
	}{
		"nominal": {
			envelope: envelope{Rates: []currencyRate{{Currency: "USD", Rate: 1.5}}},
			source: "EUR",
			target: "USD",
			err: nil,
			expected: mustParseExchangeRate(t, "1.5"),
		},

		"source = target": {
			envelope: envelope{Rates: []currencyRate{{Currency: "USD", Rate: 1}}},
			source: "EUR",
			target: "USD",
			err: nil,
			expected: mustParseExchangeRate(t, "1"),
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := tc.envelope.exchangeRate(tc.source, tc.target)
			if tc.err != err {
				t.Errorf("unable to marshal: %s", err.Error())
			}
			if got != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, got)
			}
		})
	}
}

func mustParseExchangeRate(t *testing.T, input string) money.ExchangeRate {
	t.Helper()
	rate, err := money.ParseDecimal(input)
	if err != nil {
		t.Fatalf("unable to parse exchange rate %s", input)
	}
	return money.ExchangeRate(rate)
}