package main

import (
	"fmt"
)

func main() {
	f := float64(3)
	fmt.Println(f / 4)
	var res int
	if res == 0 {
		fmt.Println("yes")
	}
}

//func findMaxAverage(nums []int, k int) float64 {
//	res := float64(-100000)
//	temp := float64(0)
//	pre := -1
//	for i, num := range nums {
//		temp += float64(num)
//		pre++
//		if i-pre > k {
//			temp -= float64(nums[pre])
//		}
//		res = max(res, temp/float64(k))
//	}
//}
