package main

import (
	"fmt"
	"sort"
)

func main() {
	a := 'a'
	b := 'b'
	fmt.Println(b - a)
}

// [l,r]
func mergeSort(nums []int, l int, r int) {
	if l < r {
		mid := (r-l)>>1 + l
		mergeSort(nums, l, mid)
		mergeSort(nums, mid+1, r)

		idx := 0
		i := l
		j := mid + 1
		help := make([]int, r-l+1)
		for i <= mid && j <= r {
			if nums[i] <= nums[j] {
				help[idx] = nums[i]
				i++
			} else {
				help[idx] = nums[j]
				j++
			}
			idx++
		}
		for i <= mid {
			help[idx] = nums[i]
			i++
			idx++
		}
		for j <= r {
			help[idx] = nums[j]
			j++
			idx++
		}
		for k, num := range help {
			nums[l+k] = num
		}
	}
	return
}

func threeSum(nums []int) [][]int {
	res := [][]int{}
	n := len(nums)
	sort.Ints(nums)
	for i, num := range nums {
		if i == len(nums)-2 {
			continue
		}
		if i > 0 && num == nums[i-1] {
			continue
		}
		if (nums[i]+nums[i+1]+nums[i+2] > 0) || (nums[i]+nums[n-1]+nums[n-2]) < 0 {
			continue
		}
		l := i + 1
		r := n - 1
		for l < r {
			temp := num + nums[l] + nums[r]
			if temp < 0 {
				l++
			} else if temp > 0 {
				r--
			} else {
				res = append(res, []int{num, nums[l], nums[r]})
				for l < r && nums[l+1] == nums[l] {
					l++
				}
				for l < r && nums[r-1] == nums[r] {
					r--
				}
				l++
				r--
			}
		}
	}
	return res
}
