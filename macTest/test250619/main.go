package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 3, 2, 3, 1}
	mergeSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
	//fmt.Println(res)
}

// []
func mergeSort(arr []int, l int, r int) int {
	if r == l {
		return 0
	}
	mid := (r-l)>>1 + l
	lCnt := mergeSort(arr, l, mid)
	rCnt := mergeSort(arr, mid+1, r)
	res := 0
	help := make([]int, r-l+1)
	idx := 0
	i := l
	j := mid + 1
	for i <= mid && j <= r {
		if arr[i] < arr[j] {
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
	for k, num := range help {
		arr[k+l] = num
	}
	return res + lCnt + rCnt
}
