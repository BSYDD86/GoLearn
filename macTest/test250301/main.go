package main

import (
	"fmt"
)

var (
	searchPath = [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	rows       int
	cols       int
)

func main() {
	num1 := "2"
	num2 := "3"
	fmt.Println(multiply(num1, num2))

}
func multiply(num1 string, num2 string) string {
	int1 := 0
	int2 := 0
	if num1 != "0" {
		for i := 0; i < len(num1); i++ {
			int1 = int1*10 + int((num1[i] - '0'))
		}
	}
	if num2 != "0" {
		for i := 0; i < len(num2); i++ {
			int2 = int2*10 + int((num2[i] - '0'))
		}
	}
	res := int1 * int2
	return string(res)
}

func canPartition(nums []int) bool {
	target := 0
	for _, num := range nums {
		target += num
	}
	if target%2 == 1 {
		return false
	}
	target /= 2
	dp := make([]int, target+1)
	for i := 0; i < target; i++ {
		for _, num := range nums {
			if i+num <= target {
				dp[i+num] = max(dp[i+num], dp[i]+num)
			}
		}
		if dp[target] == target {
			return true
		}
	}
	return false
}
