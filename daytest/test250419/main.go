package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//arr := []int{11, 13, 17, 23, 29, 31, 7, 5, 2, 3}
	//k := 3
	//thread := 5
	//println(baseNumOfSubarrays(arr, k, thread))
	i := -1
	fmt.Println(i / 2)
	str := "3 + 9 + 1 + 8 + 5 + 2 + 6"
	t := strings.Split(str, " + ")
	res := 0
	for _, c := range t {
		atoi, _ := strconv.Atoi(c)
		fmt.Println(atoi)
		res += atoi
	}
	fmt.Println(res)

	arr := []int{7, 4, 3, 9, 1, 8, 5, 2, 6}
	println(getAverages(arr, 3))
}

func getAverages(nums []int, k int) []int {
	temp := 0
	res := make([]int, len(nums))
	for i := range res {
		res[i] = -1
	}
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		temp += num
		if i-2*k-1 >= 0 {
			temp -= nums[i-k-1]
		}
		if i-2*k >= 0 {
			fmt.Println(temp)
			fmt.Println(i - k)
			res[i-k] = temp / (2*k + 1)
		}
	}
	return res
}

//1,2,3
//k=1
//i=2

func baseNumOfSubarrays(arr []int, k int, threshold int) int {
	res := 0
	for i := 0; i <= len(arr)-k; i++ {
		temp := 0
		for j := i; j < i+k; j++ {
			temp += arr[j]
		}
		if temp >= k*threshold {
			fmt.Println(i)
			fmt.Println(temp)
			res++
		}
	}
	fmt.Println("-----")
	return res
}
