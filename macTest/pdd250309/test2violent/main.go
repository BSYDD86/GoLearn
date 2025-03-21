package main

import (
	"fmt"
)

var (
	res int
)

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	res = 0
	dfs(arr, 0, 0, false)
	fmt.Println(res)
}
func dfs(arr []int, idx int, temp int, reverse bool) {
	if idx == len(arr) {
		res = max(res, temp, -temp)
		return
	}
	if reverse {
		for i := idx; i < len(arr); i++ {
			temp -= arr[i]
			res = max(res, temp, -temp)
		}
		return
	} else {
		dfs(arr, idx+1, temp+arr[idx], reverse)
		dfs(arr, idx, temp, !reverse)
	}
}

func max(a ...int) int {
	var pre int
	for _, num := range a {
		if num > pre {
			pre = num
		}
	}
	return pre
}
func min(a ...int) int {
	var pre int
	for _, num := range a {
		if num < pre {
			pre = num
		}
	}
	return pre
}
