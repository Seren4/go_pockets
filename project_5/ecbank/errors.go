package ecbank

// ecbankError defines a sentinel error.
type ecbankError string

// ecbankError implements the error interface.
func (e ecbankError) Error() string {
	return string(e)
}

const (
	ErrCallingServer = ecbankError("error calling server")
	ErrClientSide = ecbankError("error on client side")
	ErrServerSide = ecbankError("error on server side")
	ErrUnknownStatusCode = ecbankError("error unknown")

)
