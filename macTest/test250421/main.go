package main

import "fmt"

func main() {
	arr := []int{4, 3, 2, 1, 1, 2, 3, 4}
	heapSort(arr, len(arr))
	fmt.Println(arr)

}
func heapSort(nums []int, heapSize int) {
	n := len(nums)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, i, n)
	}
	for i := n - 1; i >= 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		heapify(nums, 0, i)
	}
}
func heapify(nums []int, idx int, heapSize int) {
	if idx >= heapSize {
		return
	}
	nextIdx := idx
	if 2*idx+1 < heapSize && nums[2*idx+1] > nums[nextIdx] {
		nextIdx = 2*idx + 1
	}
	if 2*idx+2 < heapSize && nums[2*idx+2] > nums[nextIdx] {
		nextIdx = 2*idx + 2
	}
	if nextIdx != idx {
		nums[idx], nums[nextIdx] = nums[nextIdx], nums[idx]
		heapify(nums, nextIdx, heapSize)
	}

}
