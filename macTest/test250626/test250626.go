package main

import "fmt"

func main() {
	m := map[int]int{}
	m[0] = 1
	v, ok := m[0]
	if ok {
		fmt.Println(v)
	}
	_, ok1 := m[1]
	if !ok1 {
		fmt.Println("no")
	}

}
