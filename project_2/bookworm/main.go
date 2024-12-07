package main

import "fmt"

func main() {
	bookworms, _ := loadBookworms("testdata/bookworms.json")
	cnt := findCommonBooks(bookworms)
	fmt.Println(cnt)
	//if err != nil {}
	fmt.Println(bookworms)



}

