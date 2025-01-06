package gordle

// corpusError defines a sentinel (recognisable) error.
type corpusError string

// Error is the implementation of the error interface corpusError
func (e corpusError) Error() string {
	return string(e)

}