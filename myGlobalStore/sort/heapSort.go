package main

import "fmt"

func maxHeap(arr []int) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
	for i := n - 1; i >= 0; i-- {
		arr[i], arr[0] = arr[0], arr[i]
		heapify(arr, i, 0)
	}
}

func heapify(arr []int, n int, i int) {
	l := 2*i + 1
	r := 2*i + 2
	largestIdx := i
	if l < n && arr[l] < arr[largestIdx] {
		largestIdx = l
	}
	if r < n && arr[r] < arr[largestIdx] {
		largestIdx = r
	}
	arr[i], arr[largestIdx] = arr[largestIdx], arr[i]
	if largestIdx != i {
		heapify(arr, n, largestIdx)
	}
}

func main() {
	arr := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	maxHeap(arr)
	fmt.Println(arr)
}
