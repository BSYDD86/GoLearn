package main

import "fmt"

func main() {
	arr := [][]int{{37, 88}, {51, 64, 65, 20, 95, 30, 26}, {9, 62, 20}, {44}}
	fmt.Println(maxValueOfCoins(arr, 2))

}

func maxValueOfCoins(piles [][]int, k int) int {
	dp := make([]int, k+1)
	for _, pile := range piles {
		cnt := 0
		for j, num := range pile {
			cnt += num

			for t := k; t >= j+1; t-- {
				dp[t] = max(dp[t], dp[t-j-1]+cnt)
			}
		}
		fmt.Println(dp)
	}
	return dp[k]

}
