package main

import (
	"math"
	"math/rand"
)

var (
	res int64
)

func main() {
	//res = 0
	//var (
	//	n int
	//)
	//fmt.Scan(&n)
	//arr := make([]int, n)
	//
	//for i := 0; i < n; i++ {
	//	fmt.Scan(&arr[i])
	//}
	//smallSum(arr, 0, n-1)
	//fmt.Println(res)
	arr := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4
	println(findKthLargest(arr, k))
}
func findKthLargest(nums []int, k int) int {
	l := 0
	r := len(nums) - 1
	for l <= r {
		if l == r {
			return nums[l]
		}
		randIdx := rand.Intn(r-l+1) + l
		nums[randIdx], nums[r] = nums[r], nums[randIdx]
		target := nums[r]

		pre := l - 1
		i := l
		j := r - 1
		for i <= j {
			if nums[i] == target {
				i++
			} else if nums[i] < target {
				pre++
				nums[i], nums[pre] = nums[pre], nums[i]
				i++

			} else {
				nums[j], nums[i] = nums[i], nums[j]
				j--
			}
		}
		nums[i], nums[r] = nums[r], nums[i]

		if r-i >= k {
			l = i + 1
		} else if k <= r-pre {
			return target
		} else {
			k -= (r - pre)
			r = pre
		}
	}
	return math.MinInt
}

func smallSum(arr []int, l int, r int) {
	if r == l {
		return
	}
	mid := (r-l)/2 + l
	smallSum(arr, l, mid)
	smallSum(arr, mid+1, r)
	merge(arr, l, mid, r)
}
func merge(arr []int, l int, mid int, r int) {
	help := make([]int, r-l+1)
	i := l
	j := mid + 1
	idx := 0

	for i <= mid && j <= r {
		if arr[i] <= arr[j] {
			res += int64((r - j + 1) * arr[i])
			help[idx] = arr[i]
			i++
		} else {
			help[idx] = arr[j]
			j++
		}
		idx++
	}
	for i <= mid {
		help[idx] = arr[i]
		i++
		idx++
	}
	for j <= r {
		help[idx] = arr[j]
		j++
		idx++
	}
	for i := 0; i < len(help); i++ {
		arr[i+l] = help[i]
	}
}
