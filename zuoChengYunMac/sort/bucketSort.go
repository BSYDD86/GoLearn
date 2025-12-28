package main

import (
	"math"
)

//优化后版本

//func main() {
//	nums := []int{13, 17, 100, 27, 25}
//	fmt.Println(nums)
//	bucketSort(nums, 0, len(nums)-1, 3)
//	fmt.Println(nums)
//}

func bucketSort(nums []int, l int, r int, digit int) {
	//Decimal
	radix := 10
	bucket := make([]int, radix)
	for d := 1; d <= digit; d++ {
		count := make([]int, radix)
		for i := l; i <= r; i++ {
			j := getDigit(nums[i], d)
			count[j]++
		}
		for i := 1; i < radix; i++ {
			count[i] = count[i] + count[i-1]
		}
		for i := r; i >= l; i-- {
			j := getDigit(nums[i], d)
			bucket[count[j]-1] = nums[i]
			count[j]--
		}
		i := l
		j := 0
		for i <= r {
			nums[i] = bucket[j]
			i++
			j++
		}
	}
}

func getDigit(num int, d int) int {
	pow := int(math.Pow(10, float64(d-1)))
	return (num / pow) % 10
}
