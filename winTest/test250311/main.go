package main

import "sort"

func main() {
	arr := []int{2, 4, 6, 4}
	sumOfBeauties(arr)
}

func sumOfBeauties(nums []int) int {
	res := 0
	n := len(nums)
	temp := nums
	sort.Ints(temp)
	for i := 1; i < n-1; i++ {
		idx := sort.SearchInts(temp, nums[i])
		idxNext := sort.SearchInts(temp, nums[i]+1)
		if idx <= i && i == idxNext-1 {
			res += 2
		} else {
			if nums[i-1] < nums[i] && nums[i] < nums[i+1] {
				res += 1
			}
		}
	}
	return res
}
