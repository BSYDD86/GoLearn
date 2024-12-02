package main

import "fmt"

func Myfunc(args interface{}) {
	fmt.Println("MyFunc running")
	fmt.Println(args)
}

type Book struct {
	name string
}

func main() {
	book := Book{name: "goland"}
	Myfunc(book)
	Myfunc(1)
	Myfunc("hello")
}
