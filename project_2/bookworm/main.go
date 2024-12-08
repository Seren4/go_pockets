package main

import "fmt"
import "os"

func main() {
	bookworms, err := loadBookworms("testdata/bookworms_bisx.json")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load bookworms: %s", err)
		os.Exit(1)
	}

	commonBooks := findCommonBooks(bookworms)

	fmt.Println("Books in common:")
	displayBooks(commonBooks)


}

