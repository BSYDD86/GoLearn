package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{3, 5, 6, 7}
	target := 9
	fmt.Println(numSubseq(nums, target))
}

var (
	res int
)

func numSubseq(nums []int, target int) int {
	sort.Ints(nums)
	j := 0
	res := 0
	const mx = 1_000_000_007
	for i, num := range nums {
		for j+1 < len(nums) {
			if num+target <= nums[j+1] {
				j++
			} else {
				break
			}
		}
		if j >= i {
			res += (1<<(j-i) - 1)
			res %= mx
		}
	}
	return res
}
