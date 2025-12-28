package main

import "fmt"

func main() {
	arr := []int{5, 2, 9, 1, 5, 6}
	insertSort(arr)
	fmt.Println(arr)
}

func insertSort(nums []int) {
	if len(nums) <= 0 {
		return
	}
	for j := 1; j < len(nums); j++ {
		num := nums[j]
		i := j - 1
		for i >= 0 && nums[i] > num {
			nums[i+1] = nums[i]
			i--
		}
		nums[i+1] = num
	}
}
