package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = dp[i-1] + arr[i-1]
	}
	fmt.Println(arr)
	fmt.Println(dp)
	var a, b int
	for {
		_, err := fmt.Scan(&a, &b)
		if err != nil {
			break
		}
		fmt.Println(dp[b+1] - dp[a])
	}
}
