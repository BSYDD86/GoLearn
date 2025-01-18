package main

import "math"

//func main() {
//	arr := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
//	println(minimumTotal(arr))
//}

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([][]int, n)
	for i, _ := range dp {
		dp[i] = make([]int, i+1)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = math.MaxInt
		}
	}
	dp[0][0] = triangle[0][0]
	for i := 0; i < n-1; i++ {
		for j := 0; j < i+1; j++ {
			dp[i+1][j] = min(dp[i+1][j], dp[i][j]+triangle[i+1][j])
			dp[i+1][j+1] = min(dp[i+1][j+1], dp[i][j]+triangle[i+1][j+1])
		}
	}
	res := math.MaxInt
	for _, num := range dp[n-1] {
		res = min(res, num)
	}
	return res
}
