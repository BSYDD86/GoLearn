package main

import "fmt"

// 最长连续字串长度，无重复字符
// 不只是数字
func main() {
	arr := "abcfghij"
	fmt.Println(getLength(arr))
}

func getLength(str string) int {
	if len(str) == 0 {
		return 0
	}
	res := 1

	m := map[rune]bool{}
	for _, c := range str {
		m[c] = true
	}
	for v := range m {
		if m[v-1] {
			continue
		}
		t := v
		cnt := 0
		for m[t] {
			cnt++
			t++
		}
		res = max(res, cnt)
	}
	return res
}

//func getLength(str string) int {
//	if len(str) == 0 {
//		return 0
//	}
//	res := 1
//	cnt := 1
//	for i := 1; i < len(str); i++ {
//		if str[i] == str[i-1]+1 {
//			cnt++
//			res = max(res, cnt)
//		} else {
//			cnt = 1
//		}
//	}
//	return res
//}
