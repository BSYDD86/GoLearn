package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	var m int
	fmt.Scan(&n)
	fmt.Scan(&m)
	arr := make([][5]int, m)
	//l,r,s,e,d
	for i := 0; i < m; i++ {
		for j := 0; j < 4; j++ {
			fmt.Scan(&arr[i][j])
		}
	}
	for i := 0; i < m; i++ {
		arr[i][4] = (arr[i][3] - arr[i][2]) / (arr[i][1] - arr[i][0])
	}
	dp := make([]int, n+1)
	for i := 0; i < m; i++ {
		l := arr[i][0]
		r := arr[i][1]
		s := arr[i][2]
		e := arr[i][3]
		d := arr[i][4]
		if 0 <= l && l < n {
			dp[l] += s
		}
		if 0 <= l+1 && l+1 < n {
			dp[l+1] += d - s
		}
		if 0 <= r+1 && r+1 <= n {
			dp[r+1] -= (d + e)
		}
		if 0 <= r+2 && r+2 <= n {
			dp[r+2] += e
		}
	}
	for i := 1; i <= n; i++ {
		dp[i] = dp[i] + dp[i-1]
	}
	for i := 1; i <= n; i++ {
		dp[i] = dp[i] + dp[i-1]
	}
	iMax := math.MinInt
	res := 0
	for i := 1; i <= n; i++ {
		temp := dp[i] - dp[i-1]
		res ^= temp
		iMax = max(iMax, temp)
	}
	fmt.Println(res, " ", iMax)
}
