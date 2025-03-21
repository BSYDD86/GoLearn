package main

import "fmt"

func main() {
	arr := []int{2, 0, 2, 1, 1, 0}
	sortColors(arr)
	fmt.Println(arr)
}

func sortColors(nums []int) {
	zeroIdx := 0
	twoIdx := len(nums) - 1
	idx := 0
	for idx < min(len(nums)-1, twoIdx) {
		for nums[idx] != 1 {
			if nums[idx] == 0 {
				nums[zeroIdx], nums[idx] = nums[idx], nums[zeroIdx]
				zeroIdx++
			} else {
				nums[twoIdx], nums[idx] = nums[idx], nums[twoIdx]
				twoIdx--
			}
		}
		idx++
	}
}
