package money

import (
	"unicode"
)

// Currency defines the code of a currency and its decimal precision.
type Currency struct {
	code string
	precision byte
}

// ErrInvalidCurrencyCode is returned when the currency to parse is not a standard 3-letter code.
const ErrInvalidCurrencyCode = Error("invalid currency code")

// ParseCurrency returns the currency associated to a name and may return ErrInvalidCurrencyCode.
func ParseCurrency(code string) (Currency, error) {
	if len(code) != 3 {
		return Currency{}, ErrInvalidCurrencyCode
	}
	// Sidequest 6.1 Make sure the currency code is made of 3 letters between A and Z. You
	// can use the regex package if you want to make things complicated, or check that each of
	// the 3 bytes is between ‘A’ and ‘Z’ included.
	// Have we properly tested everything? No

	for _, letter := range code {
		if !unicode.IsLetter(letter) {
			return Currency{}, ErrInvalidCurrencyCode
		}
	}

	switch code {
		case "IRR":
			return Currency{code: code, precision: 0}, nil
		case "CNY", "VND":
			return Currency{code: code, precision: 1}, nil
		case "BHD", "IQD", "KWD", "LYD", "OMR", "TND":
			return Currency{code: code, precision: 3}, nil
		default:
			return Currency{code: code, precision: 2}, nil
	}
}

// String implements Stringer.
func (c Currency) String() string {
	return c.code
	}

// Code provides the ISO code 
func (c Currency) Code() string {
	return c.code
	}