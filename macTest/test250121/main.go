package main

import "fmt"

func main() {
	// var a int
	// fmt.Scan(&a)
	fmt.Printf("%s", "hello world")
	intervals := [][]int{{1, 3}, {6, 9}}
	newInterval := [][]int{{2, 5}}
	res := merge(intervals, newInterval)
	fmt.Println(res)
}

// arr1 intervals
// arr2 newInterval
func merge(arr1 [][]int, arr2 [][]int) [][]int {
	res := [][]int{}
	iIdx := 0
	jIdx := 0

	for iIdx < len(arr1) || jIdx < len(arr2) {
		var (
			left  int
			right int
		)
		if iIdx == len(arr1) {
			left = arr2[jIdx][0]
			right = arr2[jIdx][1]
			jIdx++
		} else if jIdx == len(arr2) {
			left = arr1[iIdx][0]
			right = arr1[iIdx][1]
			iIdx++
		} else if arr1[iIdx][0] <= arr2[jIdx][0] {
			left = arr1[iIdx][0]
			right = arr1[iIdx][1]
			iIdx++
		} else {
			left = arr2[jIdx][0]
			right = arr2[jIdx][1]
			jIdx++
		}
		if len(res) == 0 {
			res = append(res, []int{left, right})
		} else {
			if res[len(res)-1][1] < left {
				res = append(res, []int{left, right})
			} else {
				res[len(res)-1][1] = max(res[len(res)-1][1], right)
			}
		}

	}
	return res
}
func max(i int, j int) int {
	if i >= j {
		return i
	} else {
		return j
	}
}
