package main

import (
	"fmt"
)

var (
	res float64
)

func main() {
	res = 0

	var n float64
	var k float64
	fmt.Scan(&n)
	fmt.Scan(&k)
	if n > 2*k {
		fmt.Println(-1)
	} else {
		arr := make([]int, int(n))
		for i := 0; i < int(n); i++ {
			fmt.Scan(&arr[i])
		}
		dfs(arr, int(k), 0, float64(0))
		// 使用 %.1f 格式化输出，保留一位小数
		fmt.Printf("%.1f\n", res)
	}
}

func dfs(arr []int, k int, idx int, sum float64) {
	if idx == len(arr) && k == 0 {
		if sum > res {
			res = sum
		}
		return
	}
	if idx == len(arr) || k == 0 {
		return
	}
	if idx-1+2*k < len(arr)-1 {
		return
	}
	sum += float64(arr[idx])
	dfs(arr, k-1, idx+1, sum)
	sum -= float64(arr[idx])
	if idx+1 < len(arr) {
		sum += float64((arr[idx])+arr[idx+1]) / 2.0
		dfs(arr, k-1, idx+2, sum)
	}
}
