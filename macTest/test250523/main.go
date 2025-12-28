package main

import "fmt"

func main() {
	fmt.Println(gcd(30, 20))
	fmt.Println(gcd(10, 20))
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
