package main

import "fmt"

var (
	res int
)

func main() {
	var s int
	var n int
	var m int

	fmt.Scan(&s)
	fmt.Scan(&n)
	fmt.Scan(&m)
	amount := make([]int, n)
	res = 0

	for i := 0; i < s; i++ {
		for j := 0; j < n; j++ {
			var t int
			fmt.Scan(&t)
			amount[j] = max(amount[j], 2*t)
		}
	}
	dfs(amount, 0, m, 0)
	fmt.Println(res)
}

func dfs(arr []int, idx int, money int, gain int) {
	if idx == len(arr) {
		res = max(res, gain)
		return
	}

	for i := idx; i < len(arr); i++ {
		if arr[i] <= money {
			dfs(arr, idx+1, money-arr[i], gain+2*(i+1))
		}
		dfs(arr, idx+1, money, gain)
	}
}
func max(ints ...int) int {
	var pre int
	for i := 0; i < len(ints); i++ {
		if ints[i] > pre {
			pre = ints[i]
		}
	}
	return pre
}
