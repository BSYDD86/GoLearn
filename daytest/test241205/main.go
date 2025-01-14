package main

import "fmt"

func main() {
	m := map[int]bool{1: false, 2: true}
	v, ok := m[0]
	fmt.Println(v)
	fmt.Println(ok)
}
