package money

import (
	"math"
	"fmt"
)

// ExchangeRate represents a rate to convert from a currency to another.
type ExchangeRate Decimal

type exchangeRates interface {
	FetchExchangeRate(source, target Currency) (ExchangeRate, error)
}

// Convert applies the change rate to convert an amount to a target currency.
func Convert(amount Amount, to Currency, rates exchangeRates) (Amount, error) {
	// fetch the change rate for the day
	rate, err := rates.FetchExchangeRate(amount.currency, to)        
	if err != nil {
		return Amount{}, fmt.Errorf("cannot get change rate: %w", err)
	}
	converted := applyExchangeRate(amount, to, rate)
	// Validate the converted amount is in the handled bounded range.
	if err := converted.validate(); err != nil {
		return Amount{}, err
	}
	return converted, nil
}

// applyExchangeRate returns a new Amount representing the input multiplied by the rate.
// The precision of the returned value is that of the target Currency.
// This function does not guarantee that the output amount is supported.
func applyExchangeRate(a Amount, target Currency, rate ExchangeRate) Amount {

	converted := multiply(a.quantity, rate)

	switch {
	case converted.precision > target.precision:
		converted.subunits = converted.subunits / pow10(converted.precision-target.precision)
	case converted.precision < target.precision:
		converted.subunits = converted.subunits * pow10(target.precision-converted.precision)

	}
	converted.precision = target.precision

	return Amount{quantity: converted, currency: target}

}

// multiply a Decimal with an ExchangeRate and returns the product (once simplified)
func multiply(amount Decimal, rate ExchangeRate) Decimal {
	converted := Decimal{subunits: amount.subunits * rate.subunits, precision: amount.precision + rate.precision}
	converted.simplify()
	return converted
}

// pow10 is a quick implementation of how to raise 10 to a given power.
// It's optimised for small powers, and slow for unusually high powers.
func pow10(power byte) int64 {
	switch power {
	case 0:
		return 1
	case 1:
		return 10
	case 2:
		return 100
	case 3:
		return 1000
	default:
		return int64(math.Pow(10, float64(power)))
	}
}
