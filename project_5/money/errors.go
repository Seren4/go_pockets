package money

type Error string

// Error implements the error interface.
func (e Error) Error() string {
	return string(e)
}