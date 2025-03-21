package main

import "fmt"

func main() {
	fmt.Println(returnFunc())
}

func returnFunc() (a, b int) {
	a = 10
	b = 100
	return
}
