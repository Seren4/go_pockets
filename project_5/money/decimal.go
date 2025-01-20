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
	ErrTooLarge       = Error("value over 10^12 is too big")
)

// maxDecimal value is a thousand billion, using the short scale -- 10^12.
const maxDecimal = 1e12

// ParseDecimal converts a string into its Decimal representation.
// It assumes there is up to one decimal separator, and that the separator is '.' (full stop character).
func ParseDecimal(input string) (Decimal, error) {
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
	result.simplify()

	return result, nil
}

func (d *Decimal) simplify() {
	// Using %10 returns the last digit in base 10 of a number.
	// If the precision is positive, that digit belongs to the right side of the decimal separator.
	for d.subunits%10 == 0 && d.precision > 0 {
		d.precision--
		d.subunits /= 10
	}
}

// String implements stringer and returns the Decimal formatted as
// digits and optionally a decimal point followed by digits.
func (d *Decimal) String() string {
	// Quick-win, no need to do maths.
	if d.precision == 0 {
		return fmt.Sprintf("%d", d.subunits)
	}
	centsPerUnit := pow10(d.precision)
	frac := d.subunits % centsPerUnit
	integer := d.subunits / centsPerUnit
	decimalFormat := "%d.%0" + strconv.Itoa(int(d.precision)) + "d"
	return fmt.Sprintf(decimalFormat, integer, frac)
}
