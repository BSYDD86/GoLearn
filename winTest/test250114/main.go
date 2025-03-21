package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [][]int{{1, 2}, {5, 6}, {3, 3}}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][0] < arr[j][0]
	})
	fmt.Println(arr)
}
func lemonadeChange(bills []int) bool {
	five := 0
	ten := 0
	for _, bill := range bills {
		if bill == 5 {
			five++
		} else if bill == 10 {
			if five <= 0 {
				return false
			} else {
				five--
				ten++
			}
		} else if bill == 20 {
			if ten >= 2 {
				ten -= 2
			} else if ten >= 1 && five >= 1 {
				ten--
				five--
			} else if five >= 3 {
				five -= 3
			} else {
				return false
			}
		}
	}
	return true
}
