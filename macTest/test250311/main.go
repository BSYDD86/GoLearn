package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{3, 4, -1, 1}
	fmt.Println(firstMissingPositive(arr))
}
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i]-1 >= 0 && nums[i]-1 < n && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	res := 1
	for _, num := range nums {
		if res != num {
			break
		}
		res++
	}

	return res
}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	idx := 0
	for i := range intervals {
		if intervals[i][0] <= intervals[idx][1] {
			intervals[idx][1] = max(intervals[i][1], intervals[idx][1])
		} else {
			idx++
			intervals[idx][0] = intervals[i][0]
			intervals[idx][1] = intervals[i][1]
		}
		fmt.Println(intervals)
	}
	return intervals[:idx+1]
}
