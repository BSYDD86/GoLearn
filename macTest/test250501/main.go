package main

import "fmt"

func main() {
	arr := []int{1, 1, 4, 2, 3}
	k := 5
	fmt.Println(minOperations(arr, k))
}
func minOperations(nums []int, x int) int {
	n := len(nums)
	temp := 0
	for _, num := range nums {
		temp += num
	}
	target := temp - x
	res := n + 1
	temp = 0
	pre := 0
	for i, num := range nums {
		temp += num
		for temp >= target {
			if temp == target {
				res = min(res, i-pre+1)
			}
			target -= nums[pre]
			pre++
		}
	}
	if res == n+1 {
		return -1
	}
	return res
}
