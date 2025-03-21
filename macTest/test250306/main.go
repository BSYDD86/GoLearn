package main

import "fmt"

func main() {
	s := "barfoofoobarthefoobarman"
	words := []string{"bar", "foo", "the"}
	fmt.Println(findSubstring(s, words))
}

func findSubstring(s string, words []string) []int {
	res := []int{}
	m := map[rune]int{}
	length := len(words) * len(words[0])
	for _, word := range words {
		for _, temp := range word {
			m[temp]++
		}
	}
	cur := 0
	pre := 0
	for cur < len(s) {
		m[rune(s[cur])]--
		for m[rune(s[cur])] < 0 {
			m[rune(s[pre])]++
			pre++
		}
		if cur-pre+1 == length {
			res = append(res, pre)
		}
		cur++
	}
	return res
}

func beautifulSubarrays(nums []int) (ans int64) {
	s := 0
	cnt := make(map[int]int, len(nums)+1) // 预分配空间
	cnt[0] = 1
	for _, x := range nums {
		s ^= x
		ans += int64(cnt[s])
		cnt[s]++
	}
	return
}
