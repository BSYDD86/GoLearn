package main

import (
	"fmt"
	"math"
)

func main() {
	arr := []int{1}
	arr = append(arr, 2)
	arr = append([]int{3}, arr...)
	fmt.Println(arr)
}
func myMax(nums ...int) int {
	pre := math.MaxInt32
	//var pre int
	for _, num := range nums {
		if pre > num {
			pre = num
		}
	}
	return pre
}
