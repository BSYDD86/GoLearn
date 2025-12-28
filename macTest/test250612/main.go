package main

import "fmt"

func main() {
	arr := []int{1, 2}
	fmt.Println(maxAdjacentDistance(arr))
}

func maxAdjacentDistance(nums []int) int {
	res := 0
	for i := 0; i <= len(nums); i++ {
		temp := nums[i] - nums[(i+1)%len(nums)]
		if temp < 0 {
			temp = -temp
		}
		res = max(res, temp)
	}
	return res
}
