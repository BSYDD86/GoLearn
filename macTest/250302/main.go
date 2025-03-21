package main

import "fmt"

func main() {
	arr := []int{1, 5, 11, 5}
	fmt.Println(canPartition(arr))
}
func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	sum /= 2
	fmt.Println(sum)
	dp := make([]int, sum+1)
	fmt.Println(dp)
	for _, num := range nums {
		for i := sum; i-num >= 0; i++ {
			dp[i] = max(dp[i], dp[i-num]+num)
		}
		if dp[sum] == sum {
			return true
		}
	}
	return false
}
