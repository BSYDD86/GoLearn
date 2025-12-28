package main

import "fmt"

func main() {
	str1 := "10011"
	str2 := "11001"
	fmt.Println(min(str1, str2))
}
func shortestBeautifulSubstring(s string, k int) string {
	res := ""
	cnt := 0
	pre := 0
	for i, c := range s {
		if c == '1' {
			cnt++
		}
		for cnt > k {
			if s[pre] == '1' {
				cnt--
			}
			pre++
		}
		if cnt == k {
			if res == "" {
				res = s[pre : i+1]
			} else {
				if i-pre+1 < len(res) {
					res = s[pre : i+1]
				} else if i-pre+1 == len(res) {
					res = myMin(res, s[pre:i+1])
				}
			}
		}
	}
	return res
}

func myMin(str1, str2 string) string {
	for i := range str1 {
		if str1[i] > str2[i] {
			return str1
		} else if str1[i] < str2[i] {
			return str2
		}
	}
	return str1
}
