package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 2, 3, 3, 5, 7}
	fmt.Println(minimumOperations(nums))
}

func minimumOperations(nums []int) int {
	cnt := 0
	res := 0
	m := map[int]int{}
	for _, num := range nums {
		if m[num] == 1 {
			cnt++
		}
		m[num]++
	}
	pre := 0
	for pre < len(nums) && cnt > 0 {
		res++
		for i := 0; i < 3 && pre < len(nums) && cnt > 0; i++ {
			if m[pre] == 2 {
				cnt--
			}
			m[pre]--
			pre++
		}
	}
	return res
}
