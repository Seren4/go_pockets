package money

import (
	"fmt"
	"strconv"
	"strings"
)

// Decimal is capable of storing a floating-point number with a fixed precision.
// example: 1.52 = 152 * 10^(-2) will be stored as {152, 2}
// A byte’s maximum value is 255, and we’re definitely not going to need that power of 10.
type Decimal struct {
	// subunits is the amount of subunits. Multiply it by the precision to get the real value
	subunits int64
	// Number of "subunits" in a unit, expressed as a power of 10.
	precision byte
}

const (
	ErrInvalidDecimal = Error("unable to convert the decimal")
	ErrTooLarge = Error("value over 10^12 is too big")
)

// maxDecimal value is a thousand billion, using the short scale -- 10^12.
const maxDecimal = 1e12

// ParseDecimal converts a string into its Decimal representation.
// It assumes there is up to one decimal separator, and that the separator is '.' (full stop character).
func ParseDecimal(input string) (Decimal, error){
	var result Decimal

	before, after, _ := strings.Cut(input, ".")

	units, err := strconv.ParseInt(before+after, 10, 64)
	if err != nil {
		return Decimal{}, fmt.Errorf("%w: %s", ErrInvalidDecimal, err.Error())
	}

	if units > maxDecimal {
		return Decimal{}, ErrTooLarge
	}
	precision := byte(len(after))
	result.subunits = units
	result.precision = precision

	return result, nil


}