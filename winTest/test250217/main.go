package main

import "fmt"

func main() {
	arr := []int{0}
	fmt.Println(canJump(arr))
}

func canJump(nums []int) bool {
	res := 0
	for i, n := range nums {
		if i > res {
			break
		}
		res = max(res, i+n)
		if res >= len(nums) {
			break
		}
	}
	return res >= len(nums)
}
