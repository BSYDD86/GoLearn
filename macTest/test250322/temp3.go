package main

import (
	"math/rand"
)

//	func main() {
//		nums := []int{2, 3, 4, 5, 2, 1, -1, 0}
//		quickSort(nums, 0, len(nums)-1)
//		fmt.Println(nums)
//	}
func quickSort(arr []int, l int, r int) {
	if l < r {
		pivotal := partition(arr, l, r)
		quickSort(arr, l, pivotal-1)
		quickSort(arr, pivotal+1, r)
	}
}
func partition(arr []int, l int, r int) int {
	idx := rand.Intn(r-l) + l
	arr[r], arr[idx] = arr[idx], arr[r]
	target := arr[r]
	p1 := l
	p2 := r - 1
	cur := l
	for cur <= p2 {
		if arr[cur] < target {
			arr[cur], arr[p1] = arr[p1], arr[cur]
			p1++
			cur++
		} else if arr[cur] == target {
			cur++
		} else {
			arr[cur], arr[p2] = arr[p2], arr[cur]
			p2--
		}
	}
	arr[p1], arr[r] = arr[r], arr[p1]
	return p1
}
