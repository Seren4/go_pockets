package main

import (
	"encoding/json"
	"os"
)

type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

func loadBookworms(filepath string) ([]Bookworm, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var bookworms []Bookworm
	err = json.NewDecoder(f).Decode(&bookworms)
	if err != nil {
		return nil, err
	}
	return bookworms, nil
}

// find commons books (assuming no duplication)
func findCommonBooks(bookworms []Bookworm) []Book{
	booksOnShelves := booksCount(bookworms)

	var commonBooks []Book
	
	for book, count := range booksOnShelves {
		if count > 1 {
			commonBooks = append(commonBooks, book)
		}
	}
	return commonBooks
}

func booksCount(bookworms []Bookworm) map[Book]uint{
	// map books: key = book value = uint
	count := make(map[Book]uint)
	for _, bookworm := range bookworms {
		for _, book := range bookworm.Books {
			count[book]++
		}
	}
	return count
}