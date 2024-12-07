package main

import "fmt"

func main() {
	bookworms, _ := loadBookworms("testdata/bookworms.json")
	//if err != nil {}
	fmt.Println(bookworms)



}

