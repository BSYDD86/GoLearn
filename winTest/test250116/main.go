package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{1, 2, 5}
	fmt.Println(coinChange(arr, 11))
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}
	dp[0] = 1
	for _, coin := range coins {
		for j := coin; j <= amount; j++ {
			if dp[j-coin] != math.MaxInt {
				dp[j] = min(dp[j], dp[j-coin]+1)
			}
		}
	}
	if dp[amount] == math.MaxInt {
		return -1
	}
	return dp[amount]
}
