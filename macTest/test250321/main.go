package main

import "fmt"

func main() {
	s1 := "ab"
	s2 := "eidbaooo"
	fmt.Println(checkInclusion(s1, s2))
}
func checkInclusion(s1 string, s2 string) bool {
	m := map[rune]int{}
	for _, c := range s1 {
		m[c]++
	}
	r := 0
	l := 0
	for r < len(s2) {
		c := rune(s2[r])
		m[c]--
		for m[c] < 0 && l < len(s2) {
			m[rune(s2[l])]++
			l++
		}
		if r-l+1 == len(s1) {
			return true
		}
		r++
	}
	return false
}

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
