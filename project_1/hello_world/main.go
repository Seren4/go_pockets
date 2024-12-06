package main

import "fmt"

func main() {
	greeting := greet("en")
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

/*
func greet(l language) string{
	switch l {
	case "en":
			return "Hello world"
	case "fr":
			return "Bonjour tout le monde"
	case "it":
			return "Ciao a tutti"
	case "es":
			return "Hola todos"
	default:
		return ""
		}
}
*/
