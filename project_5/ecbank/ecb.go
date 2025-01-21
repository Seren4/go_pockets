package ecbank

import (
	"learngo-pockets/moneyconverter/money"
	"net/http"
	"fmt"
	"time"
	"errors"
	"net/url"
)

// Client can call the bank to retrieve exchange rates.
type Client struct {
	client *http.Client 
	transport *http.Transport 
}

// NewBank builds a Client that can fetch exchange rates within a given timeout.
func NewClient(timeout time.Duration) Client {
	return Client{
		client: &http.Client{Timeout: timeout},
		}
	}

// FetchExchangeRate fetches the ExchangeRate for the day and returns it.
func (c Client) FetchExchangeRate(source, target money.Currency) (money.ExchangeRate, error) {

	const path = "http://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"

	resp, err := c.client.Get(path)

	if err != nil {
		var urlErr *url.Error  
    if ok := errors.As(err, &urlErr); ok && urlErr.Timeout() {
			// This is a timeout!
			return money.ExchangeRate{}, fmt.Errorf("%w: %s", ErrServerTimeOut, err)
		}
		return money.ExchangeRate{}, fmt.Errorf("%w: %s", ErrCallingServer, err)
	}
	// close the response's body
	defer resp.Body.Close()
	if err = checkStatusCode(resp.StatusCode); err != nil {
		return money.ExchangeRate{}, err
	}

	rate, err := readRateFromResponse(source.Code(), target.Code(), resp.Body)
	if err != nil {
		return money.ExchangeRate{}, err
	}
	return rate, nil

}

const (
	clientErrorClass = 4
	serverErrorClass = 5
	)
	

// checkStatusCode returns a different error depending on the returned status code.
func checkStatusCode(statusCode int) error {
	switch {

	case statusCode == http.StatusOK:
		return nil
	case httpStatusClass(statusCode) == clientErrorClass:
		return fmt.Errorf("%w: %d", ErrClientSide, statusCode)
	case httpStatusClass(statusCode) == serverErrorClass:
		return fmt.Errorf("%w: %d", ErrServerSide, statusCode)
	default:
		return fmt.Errorf("%w: %d", ErrUnknownStatusCode, statusCode)
	}
}

// httpStatusClass returns the class of a http status code.
func httpStatusClass(statusCode int) int {
	const httpErrorClassSize = 100
	return statusCode / httpErrorClassSize
}