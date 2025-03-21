package main

import (
	"fmt"
	"math"
)

func main() {
	arr := [][]int{{1, 5}}
	arrnew := []int{2, 3}
	fmt.Println(insert(arr, arrnew))

}

func insert(intervals [][]int, newInterval []int) [][]int {
	res := [][]int{}
	idx := 0
	used := false
	for idx < len(intervals) || !used {
		var pre int
		var post int
		if idx == len(intervals) {
			pre = newInterval[0]
			post = newInterval[1]
			used = true
		} else if used {
			pre = intervals[idx][0]
			post = intervals[idx][1]
			idx++
		} else {
			if intervals[idx][0] < newInterval[0] {
				pre = intervals[idx][0]
				post = intervals[idx][1]
				idx++
			} else if intervals[idx][0] > newInterval[0] {
				pre = newInterval[0]
				post = newInterval[1]
				used = true
			} else {
				pre = intervals[idx][0]
				post = max(intervals[idx][1], newInterval[1])
				idx++
				used = true
			}
		}
		if len(res) == 0 {
			res = append(res, []int{pre, post})
		} else {
			n := len(res) - 1
			if pre <= res[n][1] {
				res[n][1] = max(res[n][1], post)
			} else {
				res = append(res, []int{pre, post})
			}
		}
	}
	return res
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	m := len(nums1)
	n := len(nums2)
	l := -1
	r := m
	for l < r {
		i := (l + r) / 2
		j := (m+n+1)/2 - 2 - i
		if j+1 == n || nums1[i] < nums2[j+1] {
			l = i
		} else {
			r = i
		}
	}
	i := l
	j := (m+n+1)/2 - 2 - i
	aMin := math.MinInt
	bMin := math.MinInt
	aMax := math.MaxInt
	bMax := math.MaxInt

	if i >= 0 {
		aMin = nums1[i]
	}
	if j >= 0 {
		bMin = nums2[j]
	}
	if i+1 < m {
		aMax = nums1[i+1]
	}
	if j+1 < n {
		bMax = nums2[j+1]
	}
	if (m+n)%2 == 1 {
		return float64(max(aMin, bMin))
	} else {
		return float64(max(aMin, bMin)+min(aMax, bMax)) / 2
	}

}
