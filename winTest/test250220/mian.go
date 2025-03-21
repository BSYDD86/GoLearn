package main

import "fmt"

func main() {
	n := 50
	fmt.Println(evenOddBit(n))
}

func evenOddBit(n int) []int {
	res := make([]int, 2)
	res[0] = n
	for n != 0 {
		n = n ^ (n - 1)
		res[1]++
	}
	res[0] = res[0] - res[1]
	return res
}
