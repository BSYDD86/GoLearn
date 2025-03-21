package main

import (
	"fmt"
)

//顺序走，最多反转一次

// 4
// 1 1 -1 1
// out 3

// 5
// 1 -4 10 -30 2
// out 37
func main() {
	var n int
	fmt.Scan(&n)
	//0正
	//1负数
	//2正数一次使用
	//3负数一次使用
	//res := 0
	//dp := make([]int, 4)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	iMax := 0
	iMin := 0
	for i := 0; i < n; i++ {
		temp := arr[i]
		if temp < 0 {
			iMax, iMin = iMin, iMax
		}
		iMax = max(temp, iMax+temp)
		iMin = min(temp, iMin+temp)
		//if temp < 0 {
		//	dp[2] = dp[0] - temp
		//} else {
		//	dp[3] = dp[1] - temp
		//}
		//if temp < 0 {
		//	dp[1] += temp
		//} else {
		//	dp[0] += temp
		//}
		//res = max(res, dp[0], -dp[1], dp[2], -dp[3])
	}
	fmt.Println(max(iMax, -iMin))

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
