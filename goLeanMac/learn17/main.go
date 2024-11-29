package main

import "fmt"

type Book struct {
	name    string
	val     int
	address string
}

func main() {
	var myBook Book
	myBook.name = "ali"
	myBook.val = 1
	myBook.address = "beijing"
	fmt.Printf("%v", myBook)
}
