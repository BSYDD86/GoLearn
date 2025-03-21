package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		var m int
		monster := make([]int, n)
		demage := make([]int, m)
		cost := make([]int, m)
		for i := 0; i < n; i++ {
			fmt.Scan(monster[i])
		}
		for i := 0; i < m; i++ {
			fmt.Scan(demage)
		}
		for i := 0; i < m; i++ {
			fmt.Scan(cost)
		}

	}
}
