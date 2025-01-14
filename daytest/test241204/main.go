package main

import "fmt"

func plus(x, y int) int {
	return x + y
}

func multiPlus(x, y, z int) int {
	return x + y + z
}
func multiReturn() (int, int) {
	return 1, 3
}

func fib(i int) int {
	if i == 0 {
		return i
	}
	return i + fib(i-1)
}
func main() {
	m := map[int]int{0: 0, 2: 4}
	for k, v := range m {
		fmt.Println(k, "+", v)
	}
	v, ok := m[3]
	fmt.Println(v)
	fmt.Println(ok)

}
