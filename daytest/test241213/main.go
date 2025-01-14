package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{1, 2, 3}
	swap(arr, 0, 2)
	fmt.Printf("aaa")
}
func swap(arr []int, l int, r int) {
	arr[l], arr[r] = arr[r], arr[l]
}

func canJump(nums []int) bool {
	n := len(nums)
	dp := make([]bool, n)
	dp[0] = true
	for i, n := range nums {
		if dp[i] == false {
			return false
		}
		t := i + n
		if t >= n-1 {
			return true
		}
		for j := i + 1; j <= t; j++ {
			dp[j] = true
		}
	}
	return false
}
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i, _ := range dp {
		dp[i] = math.MaxInt
	}
	dp[0] = 0
	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]

}
