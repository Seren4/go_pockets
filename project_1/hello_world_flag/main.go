package main

import (
	"flag"
	"fmt"
)

func main() {
	var lang string
	flag.StringVar(&lang, "lang", "en", "The required language, e.g. en, it, fr, ...") 
	flag.Parse()
	greeting := greet(language(lang))
	
	//OR
	/*
	lang := flag.String("lang", "en", "The required language, e.g. en, it, fr, ...") 
	flag.Parse()
	greeting := greet(language(*lang))
	*/

	// todo add validation on lang
	fmt.Println(greeting)
}

type language string 

var phrasebook = map[language]string{
	"en": "Hello world",
	"fr": "Bonjour tout le monde",
	"it": "Ciao a tutti",
	"es": "Hola todos",
}

func greet(l language) string{
	greeting, ok := phrasebook[l]
	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}
	return greeting
}
