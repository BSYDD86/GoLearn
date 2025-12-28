package main

import (
	"fmt"
)

func main() {
	arr := []int{5, 4, 2, 4}
	fmt.Println(continuousSubarrays(arr))
}

func numSubarraysWithSum(nums []int, goal int) int {
	sum1 := 0
	sum2 := 0
	l1 := 0
	l2 := 0
	res := 0
	for r, num := range nums {
		sum1 += num
		for l1 < r && sum1 > goal {
			sum1 -= nums[l1]
			l1++
		}
		sum2 += num
		for l2 < r && sum2 >= goal {
			sum2 -= nums[l2]
			l2++
		}
		res += (l2 - l1)
	}
	return res
}

func continuousSubarrays(nums []int) int64 {
	minQ := []int{}
	maxQ := []int{}
	l := 0
	res := 0
	const mx = 2
	for r, num := range nums {
		for len(minQ) > 0 && num <= minQ[len(minQ)-1] {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, r)
		for len(maxQ) > 0 && num >= maxQ[len(maxQ)-1] {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, r)
		for maxQ[0]-minQ[0] > mx {
			l++
			if maxQ[0] < l {
				maxQ = maxQ[1:]
			}
			if minQ[0] < l {
				minQ = minQ[1:]
			}
		}
		if r == 2 {
			fmt.Println(minQ)
			fmt.Println(maxQ)
		}
		fmt.Println(l, "---", r)
		res += (r - l + 1)
	}
	return int64(res)
}
func longestSubsequenceRepeatedK(s string, k int) string {
	m := map[string]int{}
	for i := 0; i < len(s); i++ {
		mcnt := 0
		for j := i + 1; j <= len(s); j++ {
			subArr := s[i:j]
			mcnt++
			m[subArr]++
		}
	}
	var res string
	cnt := 0
	for key, v := range m {
		cnt += v
		if v == k {
			res = max(res, key)
		}
	}
	fmt.Println(cnt)
	return res
}
