package ecbank

import "learngo-pockets/moneyconverter/money"

type envelope struct {
	Rates []currencyRate `xml:"Cube>Cube>Cube"`
}

type currencyRate struct {
	Currency string             `xml:"currency,attr"`
	Rate     money.ExchangeRate `xml:"rate,attr"`
}

const baseCurrencyCode = "EUR"

// exchangeRates builds a map of all the supported exchange rates.
func (e envelope) exchangeRates() map[string]money.ExchangeRate {

	rates := make(map[string]money.ExchangeRate, len(e.Rates)+1)

	for _, c := range e.Rates {
		rates[c.Currency] = c.Rate
	}
	rates[baseCurrencyCode] = 1
	return rates
}

// exchangeRate reads the change rate from the Envelope's contents.
func (e envelope) exchangeRate(source, target string) (money.ExchangeRate, error) {
	if source == target {
		return money.ExchangeRate{subunits: 1, precision: 0}, nil
	}
	rates := e.exchangeRates()
	//rates := e.mappedChangeRates()
	from := rates[source]
	to := rates[target]
}
