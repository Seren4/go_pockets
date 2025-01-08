package money_test

import (
	"learngo-pockets/moneyconverter/money"
	"reflect"
	"testing"
)

func mustParseCurrency(t *testing.T, code string) money.Currency {
	t.Helper()
	currency, err := money.ParseCurrency(code)
	if err != nil {
		// we are not using t.Fail but t.Fatal, which stops the test run immediately.
		t.Fatalf("cannot parse currency %s code", code)
	}
	return currency
}

func mustParseAmount(t *testing.T, value string, code string) money.Amount {
	t.Helper()
	decimal, err := money.ParseDecimal(value)
	if err != nil {
		t.Fatalf("invalid number %s", value)
	}
	currency, err := money.ParseCurrency(code)
	if err != nil {
		t.Fatalf("Invalid currency code %s", code)
	}

	amount, err := money.NewAmount(decimal, currency)
	if err != nil {
		t.Fatalf("cannot create amount with value %v and currency code %s", decimal, code)
	}

	return amount
}

func TestConvert(t *testing.T) {
	tt := map[string]struct {
		amount money.Amount
		to money.Currency
		validate func(t *testing.T, got money.Amount, err error)
	}{
		"34.98 USD to EUR": {
			amount: mustParseAmount(t, "34.98", "USD"),
			to: mustParseCurrency(t, "EUR"),
			validate: func(t *testing.T, got money.Amount, err error) {
				if err != nil {
					t.Errorf("expected no error, got %s", err.Error())
				}
				expected := mustParseAmount(t, "69.96", "EUR")
				if !reflect.DeepEqual(got, expected) {
					t.Errorf("expected %v, got %v", expected, got)
				}
			},
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := money.Convert(tc.amount, tc.to)
			tc.validate(t, got, err)
		})
			}
		}
	