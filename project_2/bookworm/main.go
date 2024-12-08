package main

import (
	"fmt"
	"os"
	"flag"
)

func main() {
	var filePath string
	flag.StringVar(&filePath, "filePath", "testdata/bookworms.json", "The bookworms filepath") 
	flag.Parse()
	bookworms, err := loadBookworms(filePath)
	
	//bookworms, err := loadBookworms("testdata/bookworms.json")
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load bookworms: %s", err)
		os.Exit(1)
	}

	commonBooks := findCommonBooks(bookworms)

	fmt.Println("Books in common:")
	displayBooks(commonBooks)


}

