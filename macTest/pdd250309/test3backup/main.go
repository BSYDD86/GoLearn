package main

import (
	"fmt"
)

//读书翻倍速度

// 4 1
// 1 2 3 4
// out -1

// 5 3
// 1 2 3 2 1
// out 6.0

func main() {
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
		split := int(n - k)
		res := float64(0)
		for i := 0; i < split; i += 2 {
			res += float64(arr[i]+arr[i+1]) / 2.0
		}
		for i := split; i < int(n); i++ {
			res += float64(arr[i])
		}
		fmt.Println(res)
	}
}

func dfs() {

}
