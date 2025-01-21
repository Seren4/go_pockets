package money_test

import (
	"learngo-pockets/moneyconverter/money"
	"reflect"
	"testing"
	"errors"
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

func mustParseExchangeRate(t *testing.T, input string) money.ExchangeRate {
	t.Helper()
	rate, err := money.ParseDecimal(input)
	if err != nil {
		t.Fatalf("unable to parse exchange rate %s", input)
	}
	return money.ExchangeRate(rate)
}

// stubRate is a very simple stub for the exchangeRates.
type stubRate struct {
	rate money.ExchangeRate
	err error 
}

// FetchExchangeRate implements the interface exchangeRates with the same signature but fields are unused for tests purposes.
func (m stubRate) FetchExchangeRate(_, _ money.Currency) (money.ExchangeRate, error) {
    return m.rate, m.err
}

func TestConvert(t *testing.T) {
	tt := map[string]struct {
		amount money.Amount
		to money.Currency
		stub stubRate
		validate func(t *testing.T, got money.Amount, err error)
	}{
		"34.98 USD to EUR": {
			amount: mustParseAmount(t, "34.98", "USD"),
			to: mustParseCurrency(t, "EUR"),
			stub: stubRate{rate: mustParseExchangeRate(t, "2.00"), err: nil},
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
		"error on finale validatio": {
			amount: mustParseAmount(t, "500000000001", "USD"),
			to: mustParseCurrency(t, "EUR"),
			stub: stubRate{rate: money.ExchangeRate{}, err: money.ErrTooLarge},
			validate: func(t *testing.T, got money.Amount, err error) {
				if !errors.Is(err, money.ErrTooLarge){
					t.Errorf("expected error %s, got %s", money.ErrTooLarge, err.Error())
				}
			},
		},
	}
	for name, tc := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := money.Convert(tc.amount, tc.to, tc.stub)
			tc.validate(t, got, err)
		})
			}
		}
	