package main

import "fmt"

func main() {
	m := map[int]int{}
	m[1] = 1
	v1 := m[1]
	v2 := m[2]
	fmt.Println(v1)
	fmt.Println(v2)

}
