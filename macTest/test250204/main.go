package main

import "fmt"

func main() {
	res := "cbacdcbc"
	fmt.Println(removeDuplicateLetters(res))
}
func removeDuplicateLetters(s string) string {
	m := map[rune]bool{}
	cnt := 0
	res := s
	for _, c := range s {
		if !m[c] {
			cnt++
		}
		m[c] = true
	}
	pre := 0
	mRes := map[rune]int{}
	for i, c := range s {
		if mRes[c] == 0 {
			cnt--
		}
		mRes[c]++
		for mRes[c] > 1 {
			if mRes[rune(s[pre])] == 1 {
				cnt++
			}
			mRes[rune(s[pre])]--
			pre++
		}
		if cnt == 0 {
			res = min(res, s[pre:i+1])
		}
	}
	return res

}

func removeKdigits(num string, k int) string {
	st := []rune{}
	for _, n := range num {
		for len(st) > 0 && n < st[len(st)-1] && k > 0 {
			st = st[:len(st)-1]
			k--
		}

		st = append(st, n)
	}
	st = st[:max(len(st)-k, 0)]
	i := 0
	for i < len(st) {
		if st[i] == '0' {
			i++
		} else {
			break
		}
	}
	if len(st) == 0 {
		return "0"
	}
	return string(st[i:])

}

func sortArrayByParityII(nums []int) []int {
	oddIdx := 1
	for i := range nums {
		if i%2 == 1 {
			for oddIdx < len(nums) && nums[i]%2 == 1 {
				nums[i], nums[oddIdx] = nums[oddIdx], nums[i]
				oddIdx += 2
			}
		}
	}
	return nums
}
