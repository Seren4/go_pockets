package ecbank

import (
	"testing"
	"fmt"
	"net/http/httptest"
	"net/http"
	"learngo-pockets/moneyconverter/money"
	"errors"
)

func TestEuroCentralBank_FetchExchangeRate_Success(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `<?xml version="1.0" encoding="UTF-8"?><gesmes:Envelope>
		<gesmes:subject>Reference rates</gesmes:subject>
		<Cube time='2025-01-17'><Cube>
		<Cube currency='USD' rate='2'/>
		<Cube currency='RON' rate='6'/>
		</Cube></Cube></gesmes:Envelope>`)
	}))
	defer ts.Close() 

	ecb := Client{
		url: ts.URL,
	}
	got, err := ecb.FetchExchangeRate(mustParseCurrency(t, "USD"), mustParseCurrency(t,"RON")) 
	want := mustParseDecimal(t, "3")

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if want != money.Decimal(got) {
		t.Errorf("FetchExchangeRate got = %v, want %v", money.Decimal(got), want)
	}
}

func TestEuroCentralBank_FetchExchangeRate_Error500(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer ts.Close() 

	ecb := Client{
		url: ts.URL,
	}
	got, err := ecb.FetchExchangeRate(mustParseCurrency(t, "USD"), mustParseCurrency(t,"RON")) 
	want := mustParseDecimal(t, "0")

	if !errors.Is(err, ErrServerSide){
		t.Errorf("unexpected error: %v, want: %v", err, ErrServerSide)
	}
	
	if want != money.Decimal(got) {
		t.Errorf("FetchExchangeRate got = %v, want %v", money.Decimal(got), want)
	}
}

func TestEuroCentralBank_FetchExchangeRate_ErrorXML(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `<?xml version="1.0" encoding="UTF-8"?><gesmes:Envelope>
		<gesmes:subject>Reference rates</gesmes:subject>
		<Cube time='2025-01-17'><Cube>
		<Cube time='2025-01-17'><Cube>
		<Cube currency='USD' rate='2'/>
		<Cube currency='RON' rate='6'/>
		</Cube></Cube></gesmes:Envelope>`)
	}))
	defer ts.Close() 

	ecb := Client{
		url: ts.URL,
	}
	got, err := ecb.FetchExchangeRate(mustParseCurrency(t, "USD"), mustParseCurrency(t,"RON")) 
	want := mustParseDecimal(t, "0")

	if !errors.Is(err, ErrUnexpectedFormat){
		t.Errorf("unexpected error: %v, want: %v", err, ErrServerSide)
	}
	
	if want != money.Decimal(got) {
		t.Errorf("FetchExchangeRate got = %v, want %v", money.Decimal(got), want)
	}
}

func TestEuroCentralBank_FetchExchangeRate_CurrencyNotFound(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, `<?xml version="1.0" encoding="UTF-8"?><gesmes:Envelope>
		<gesmes:subject>Reference rates</gesmes:subject>
		<Cube time='2025-01-17'><Cube>
		<Cube currency='USD' rate='2'/>
		<Cube currency='RON' rate='6'/>
		</Cube></Cube></gesmes:Envelope>`)
	}))
	defer ts.Close() 

	ecb := Client{
		url: ts.URL,
	}
	got, err := ecb.FetchExchangeRate(mustParseCurrency(t, "USA"), mustParseCurrency(t,"ROI")) 
	want := mustParseDecimal(t, "0")

	if !errors.Is(err, ErrChangeRateNotFound){
		t.Errorf("unexpected error: %v, want: %v", err, ErrServerSide)
	}
	
	if want != money.Decimal(got) {
		t.Errorf("FetchExchangeRate got = %v, want %v", money.Decimal(got), want)
	}
}

func mustParseCurrency(t *testing.T, input string) money.Currency {
	t.Helper()
	curr, err := money.ParseCurrency(input)
	if err != nil {
		t.Fatalf("unable to parse currency rate %s", input)
	}
	return curr
}
func mustParseDecimal(t *testing.T, input string) money.Decimal {
	t.Helper()
	decimal, err := money.ParseDecimal(input)
	if err != nil {
		t.Fatalf("unable to parse decimal %s", input)
	}
	return decimal
}