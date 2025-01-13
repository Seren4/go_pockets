package money

// Amount defines a quantity of money in a given Currency.
type Amount struct {
	quantity Decimal
	currency Currency
}

// ErrTooPrecise is returned if the number is too precise for the currency.
const ErrDecimalTooPrecise = Error("quantity is too precise")

func NewAmount(quantity Decimal, currency Currency) (Amount, error) {
	if quantity.precision > currency.precision {
		return Amount{}, ErrDecimalTooPrecise
	}
	quantity.precision = currency.precision
	return Amount{quantity: quantity, currency: currency}, nil
}

// validate returns an error if and only if an Amount is unsafe to use.
func (a Amount) validate() error{
	switch {
	case a.quantity.subunits > maxDecimal:
		return ErrTooLarge
	case a.quantity.precision > a.currency.precision:
		return ErrDecimalTooPrecise
	}
	return nil
}

func (a Amount) String() string {
	return a.quantity.String() + " " + a.currency.code
	}