package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 3}
	fmt.Println(sort.SearchInts(arr, 1))
	fmt.Println(sort.SearchInts(arr, 2))
	fmt.Println(sort.SearchInts(arr, 3))
}

// []
func mergeAndCal(arr []int, l int, r int) int {
	if l == r {
		return 0
	}
	mid := (r-l)>>1 + l
	lCnt := mergeAndCal(arr, l, mid)
	rCnt := mergeAndCal(arr, mid+1, r)
	help := make([]int, r-l+1)
	i := l
	j := mid + 1
	idx := 0
	res := 0
	for i <= mid && j <= r {
		if arr[i] <= arr[j] {
			help[idx] = arr[i]
			res += (r - j + 1) * arr[i]
			i++
		} else {
			help[idx] = arr[j]
			j++
		}
		idx++
	}
	for i <= mid {
		help[idx] = arr[i]

		idx++
		i++
	}
	for j <= r {
		help[idx] = arr[j]
		idx++
		j++
	}
	for i := range help {
		arr[i+l] = help[i]
	}
	return res + lCnt + rCnt
}

func divideArray(nums []int, k int) [][]int {
	sort.Ints(nums)
	n := len(nums)
	res := [][]int{}
	for i := 0; i < n; i += 3 {
		if nums[i]+k > nums[i+2] {
			return [][]int{}
		} else {
			temp := []int{nums[i], nums[i+1], nums[i+2]}
			res = append(res, temp)
		}
	}
	return res
}
