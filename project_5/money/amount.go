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