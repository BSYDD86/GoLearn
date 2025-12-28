package main

import "fmt"

func main() {
	arr := []int{1, 7}
	fmt.Println(maximumPrimeDifference(arr))
}

func maximumPrimeDifference(nums []int) int {
	pre := -1
	post := -1
	for i, num := range nums {
		fmt.Println(num)
		fmt.Println(isPrime(num))
		if isPrime(num) {
			if pre == -1 {
				pre = i
			}
			post = i
		}
	}
	fmt.Println(pre)
	fmt.Println(post)

	return post - pre
}

func isPrime(num int) bool {
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
