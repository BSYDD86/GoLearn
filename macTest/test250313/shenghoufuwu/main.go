package main

import (
	"fmt"
)

func main() {
	arr := []int{9, 2, 5, 3, 7, 101, 18}
	fmt.Println(getIncrLen(arr))
}

func getIncrLen(nums []int) int {
	arr := make([]int, len(nums))
	res := 0
	for _, num := range nums {
		l := 0
		r := res
		for l < r {
			mid := (l + r) / 2
			if arr[mid] < num {
				l = mid + 1
			} else {
				r = mid
			}
		}
		arr[l] = num
		if r == res {
			res++
		}
	}
	fmt.Println(arr)
	return res
}
