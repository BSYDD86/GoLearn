package main

import "fmt"

func main() {
	arr := []int{4, 5, 6, 7, 0, 1, 2}
	fmt.Println(findMin(arr))
}

func findMin(nums []int) int {
	n := len(nums)
	target := nums[n-1]
	l := 0
	r := n - 2
	//[]
	for l <= r {
		mid := (r + l) / 2
		if nums[mid] > target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return nums[l]
}

func evenOddBit(n int) []int {
	even := 0
	odd := 0
	for n != 0 {
		if n&1 == 1 {
			odd++
		} else {
			even++
		}
		n = n >> 1
	}
	return []int{even, odd}
}
