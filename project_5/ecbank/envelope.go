package ecbank

import (
	"learngo-pockets/moneyconverter/money"
	"io"
	"fmt"
	"encoding/xml"
)

type envelope struct {
	Rates []currencyRate `xml:"Cube>Cube>Cube"`
}

type currencyRate struct {
	Currency string             `xml:"currency,attr"`
	Rate     float64						`xml:"rate,attr"`
}

const baseCurrencyCode = "EUR"

// exchangeRates builds a map of all the supported exchange rates.
func (e envelope) mappedChangeRates() map[string]float64 {

	rates := make(map[string]float64, len(e.Rates)+1)

	for _, c := range e.Rates {
		rates[c.Currency] = c.Rate
	}
	rates[baseCurrencyCode] = 1.
	return rates
}

// exchangeRate reads the change rate from the Envelope's contents.
func (e envelope) exchangeRate(source, target string) (money.ExchangeRate, error) {

	if source == target {
		one, err := money.ParseDecimal("1")
		if err != nil {
			return money.ExchangeRate{}, fmt.Errorf("unable to create a rate of value 1: %w", err)
		}
		return money.ExchangeRate(one), nil
	}
	rates := e.mappedChangeRates()
	
	from, fromFound := rates[source]
	if !fromFound {
		return money.ExchangeRate{}, fmt.Errorf("failed to find the source currency %s", source)
	}
	to, toFound := rates[target]
	if !toFound {
		return money.ExchangeRate{}, fmt.Errorf("failed to find the target currency %s", target)
	}

	// Use a precision of 10 digits after the decimal separator.
	// This should be enough, as most currencies only use 5 digits
	rate, err := money.ParseDecimal(fmt.Sprintf("%.10f", to/from))

	if err != nil {
		return money.ExchangeRate{}, fmt.Errorf("unable to parse exchange rate from %s to %s: %w", source, target, err)
	}

	return money.ExchangeRate(rate), nil
}


func readRateFromResponse(source, target string, respBody io.Reader) (money.ExchangeRate, error) {
	// read the response
	decoder := xml.NewDecoder(respBody)
	var ecbMessage envelope
	err := decoder.Decode(&ecbMessage)
	if err != nil {
		return money.ExchangeRate{}, fmt.Errorf("%w: %s", ErrUnexpectedFormat, err)
	}
	rate, err := ecbMessage.exchangeRate(source, target)
	if err != nil {
		return money.ExchangeRate{}, fmt.Errorf("%w: %s", ErrChangeRateNotFound, err)
	}
	return rate, nil
}