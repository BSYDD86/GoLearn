package main

import "fmt"

func main() {
	arr := []int{-2, 3, -4}
	fmt.Println(maxProduct(arr))
}
func maxProduct(nums []int) int {
	iMax := 1
	iMin := 1
	res := nums[0]
	for _, num := range nums {
		if num < 0 {
			iMax, iMin = iMin, iMax
		}
		iMax = max(iMin*num, num)
		iMin = min(iMin*num, num)
		res = max(res, iMax)
	}
	return res
}
