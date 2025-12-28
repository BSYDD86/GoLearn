package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	l := "2"
	r := "7"
	b := 2
	fmt.Println(countNumbers(l, r, b))
}

//
//func smallestPalindrome(s string) string {
//	n := len(s)
//	arr := []byte(s[:n/2])
//	slices.Sort(arr)
//	res := string(arr)
//	if n%2 == 1 {
//		res += string(s[n/2])
//	}
//	slices.Reverse(arr)
//	return res + string(arr)
//}

var (
	res int
)

func countNumbers(l string, r string, b int) int {
	lArr := []int{}
	rArr := []int{}
	lInt, _ := strconv.Atoi(l)
	tempLInt := lInt
	rInt, _ := strconv.Atoi(r)
	tempRInt := rInt
	for lInt > 0 {
		lArr = append(lArr, lInt%b)
		lInt /= b
	}
	for rInt > 0 {
		rArr = append(rArr, rInt%b)
		rInt /= b
	}
	for i := lInt; i <= tempRInt-tempLInt; i++ {
		if check(lArr) {
			res++
		}
		lArr = myArrAdd(lArr, 0, b)
	}
	return res
}

func myArrAdd(nums []int, base int, scale int) []int {
	for {
		if base == len(nums) {
			nums = append(nums, 1)

			break
		}
		if nums[base] < scale-1 {
			nums[base]++

			break
		} else {
			nums[base] = 0
			base++
		}
	}
	return nums
}
func check(nums []int) bool {
	isSame := true
	for i := 0; i < len(nums)-1; i++ {
		fmt.Println(nums)
		if nums[i] != nums[i+1] {
			isSame = false
		}
		if nums[i] > nums[i+1] {
			fmt.Println(nums, "true")
			return true
		}
	}
	if isSame {
		fmt.Println(nums, "true")
		return true
	}
	return false
}

func myStrMin(str1 string, str2 string) string {
	i := 0
	j := 0
	for i < len(str1) && j < len(str2) {
		if str1[i] < str2[j] {
			return str1
		} else if str1[i] == str2[j] {
			i++
			j++
		} else {
			return str2
		}
	}
	if len(str1) < len(str2) {
		return str2
	} else {
		return str1
	}
}

func smallestPalindrome(s string) string {
	arr := make([]int, 26)
	for _, c := range s {
		arr[c-'a']++
	}
	var (
		temp string
		res  string
	)
	for i, c := range arr {
		if c%2 == 1 {
			temp = string('a' + i)
		}
		repeat := strings.Repeat(string(i+'a'), c/2)
		res += repeat
	}
	cnt := len(res)
	if temp != "" {
		res += temp
	}
	for i := cnt - 1; i >= 0; i-- {
		res += string(res[i])
	}
	return res
}

func findClosest(x int, y int, z int) int {
	cntX := 0
	cntY := 0
	if z > x {
		cntX = z - x
	} else {
		cntX = x - z
	}
	if z > y {
		cntY = z - y
	} else {
		cntY = y - z
	}
	if cntX < cntY {
		return 1
	} else if cntX == cntY {
		return 0
	} else {
		return 2
	}
}
