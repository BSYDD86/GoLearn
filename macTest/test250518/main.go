package main

import "fmt"

func main() {
	arr := []int{1, 3, 2}
	fmt.Println(smallestIndex(arr))
}

func smallestIndex(nums []int) int {
	res := -1
	for i, num := range nums {
		if i == getBit(num) {
			res = i
		}
	}
	return res
}

func getBit(num int) int {
	res := 0
	for num > 0 {
		res += (num % 10)
		num /= 10
	}
	return res
}
