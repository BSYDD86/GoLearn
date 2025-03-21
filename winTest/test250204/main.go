package main

import (
	"fmt"
	"strconv"
)

func main() {
	str1 := "02"
	str2 := "002"
	num1, _ := strconv.Atoi(str1)
	num2, _ := strconv.Atoi(str2)
	fmt.Println(num1 == num2)
}

func findRepeatedDnaSequences(s string) []string {
	res := []string{}
	m := map[string]int{}
	for i := 0; i < len(s)-10; i++ {
		for j := i + 10; j <= len(s); j++ {
			temp := s[i:j]
			if m[temp] == 1 {
				res = append(res, temp)
			}
			m[temp]++
		}
	}
	return res
}

func isPalindrome(s string) bool {
	res := []rune{}
	for _, c := range s {
		if 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z' {
			res = append(res, c)
		}
	}
	l := 0
	r := len(res) - 1
	fmt.Println(string(res))
	for l < r {
		if !(res[l] == res[r] || res[r]-res[l] == 32 || res[l]-res[r] == 32) {
			fmt.Println(l)
			fmt.Println(r)
			fmt.Println(len(res))
			fmt.Println(string((res)))
			return false
		}
		l++
		r--
	}
	return true
}
func merge(nums1 []int, m int, nums2 []int, n int) {
	i := 0
	j := 0
	res := []int{}
	for i < m || j < n {
		var t int
		if i == m {
			t = nums2[j]
			j++
		} else if j == n {
			t = nums1[i]
			i++
		} else if nums1[i] <= nums2[j] {
			t = nums1[i]
			i++
		} else {
			t = nums2[j]
			j++
		}
		res = append(res, t)
	}
	nums1 = res
}
