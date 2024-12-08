package main

import (
	"testing"
	"reflect"
)


var (
	handmaidsTale = Book{Author: "Margaret Atwood", Title: "The Handmaid's Tale"}
	oryxAndCrake = Book{Author: "Margaret Atwood", Title: "Oryx and Crake"}
	theBellJar = Book{Author: "Sylvia Plath", Title: "The Bell Jar"}
	janeEyre = Book{Author: "Charlotte Brontë", Title: "Jane Eyre"}
)

func TestLoadBookworms(t *testing.T) {
	type testCase struct {
		bookwormsFile	string
		want []Bookworm
		wantErr	bool
	}

	var tests = map[string]testCase{
		"file exists": {
			bookwormsFile: "testdata/bookworms.json",
			want: []Bookworm{
				{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
				{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
			},
			wantErr: false,
		},
		"file doesn't exist": {
			bookwormsFile: "testdata/no_file_here.json",
			want:          nil,
			wantErr:       true,
		},
		"invalid JSON": {
			bookwormsFile: "testdata/invalid.json",
			want:          nil,
			wantErr:       true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := loadBookworms(tc.bookwormsFile)

			if err != nil && !tc.wantErr {
				t.Fatalf("expected an error %s, got none", err.Error())
			}
			if err == nil && tc.wantErr {
				t.Fatalf("expected no error, got one %s", err.Error())
			}
			/*
			if !equalBookworms(got, tc.want){
				t.Fatalf("different result: got %v, expected %v", got, tc.want)
			}
			*/
			// OR compare with reflect
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("different result: got %v, expected %v", got, tc.want)
			}
		})
	}
}

func equalBookworms(bookworms, target []Bookworm) bool{
	if len(bookworms) != len(target) {
		return false
	}
	for i := range bookworms {
		if bookworms[i].Name != target[i].Name {
			return false
		}
		if !equalBooks(bookworms[i].Books, target[i].Books) {
			return false
		}
	}
	return true
}

func equalBooks(books, target []Book) bool{
	if len(books) != len(target) {
		return false
	}
	for i := range books {
		if books[i] != target[i] {
			return false
		}
	}
	return true
	}

	func equalBooksCount(got, want map[Book]uint) bool{
	if len(got) != len(want) {
		return false
	}
	for book, targetCount := range want {
		count, ok := got[book]
		if !ok || targetCount != count {
			return false
		}
	}
	return true
	}

	func TestBooksCount(t *testing.T) {
		tt := map[string]struct{
			input []Bookworm
			want map[Book]uint
		}{
			"nominal use case": {
				input: []Bookworm{
					{Name: "Fadi", Books: []Book{handmaidsTale, theBellJar}},
					{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
				},
				want: map[Book]uint{janeEyre:1, oryxAndCrake:1, handmaidsTale:2, theBellJar:1},
			},
			"no input": {
				input: []Bookworm{},
				want: map[Book]uint{},
			},
			"no books": {
				input: []Bookworm{
					{Name: "Fadi", Books: []Book{}},
					{Name: "Peggy", Books: []Book{}},
				},
				want: map[Book]uint{},
			},
			"with book duplication": {
				input: []Bookworm{
					{Name: "Fadi", Books: []Book{handmaidsTale, handmaidsTale, theBellJar}},
					{Name: "Peggy", Books: []Book{oryxAndCrake, handmaidsTale, janeEyre}},
				},
				want: map[Book]uint{janeEyre:1, oryxAndCrake:1, handmaidsTale:3, theBellJar:1},
			},
		}
		for name, tc := range tt {
			t.Run(name, func(t *testing.T) {
				got := booksCount(tc.input)
				if !reflect.DeepEqual(got, tc.want) {
					t.Fatalf("different result: got %v, expected %v", got, tc.want)
				}
			})
		}
	}