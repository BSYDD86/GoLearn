package main

import "fmt"

func main() {
	s := "(())"
	for _, c := range s {
		fmt.Println(c)
	}
}
