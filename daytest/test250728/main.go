package main

import "fmt"

func main() {
	m := map[int]int{}
	m[0] = 0
	delete(m, 0)
	fmt.Println(len(m))
}
