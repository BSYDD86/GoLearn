package main

import "fmt"

func main() {
	nums := []int{5, 7, 7, 8, 8, 10}
	fmt.Println(getIdx(nums, 8))

}

func getIdx(nums []int, target int) int {
	l := 0
	r := len(nums)
	for l < r {
		mid := (r-l)/2 + l
		if nums[l] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}

// []
func search1(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	l := 0
	r := len(nums) - 1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	fmt.Println("l is", l)
	fmt.Println(r)
	return l
}
