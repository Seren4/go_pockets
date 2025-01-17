package main

import (
	"flag"
	"fmt"
	"learngo-pockets/moneyconverter/money"
	"os"
)

func main() {
	from := flag.String("from", "", "source currency, required")
	to := flag.String("to", "EUR", "target currency, required")
	flag.Parse()
	value := flag.Arg(0)
	if value == "" {
		_, _ = fmt.Fprintln(os.Stderr, "missing amount to convert")
		flag.Usage()
		os.Exit(1)
	}

	fromCurrency, err := money.ParseCurrency(*from)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to parse source currency %q: %s.\n", *from, err)
	}
	toCurrency, err := money.ParseCurrency(*to)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to parse target currency %q: %s.\n", *to, err)
	}

	quantity, err := money.ParseDecimal(value)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "unable to parse value %q: %s.\n", value, err)
	}

	amount, err := money.NewAmount(quantity, fromCurrency)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	result, err := money.Convert(amount, toCurrency)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
	fmt.Println("Amount:", amount, "; To Currency:", toCurrency)
	fmt.Println(result)




}
