package main

import "fmt"

func main() {
	s := "a"
	dict := []string{"a"}
	fmt.Println(wordBreak(s, dict))

	//
	//a := "aa"
	////b := "bb"
	//c := "aabb"
	//for _, t := range c {
	//	fmt.Println(t == a[0])
	//}

}
func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	dp := make([]bool, n+1)
	m := map[string]bool{}
	for _, c := range wordDict {
		m[c] = true
	}
	dp[0] = true
	for i := 0; i <= n; i++ {
		if !dp[i] {
			continue
		}
		for j := i + 1; j < n; j++ {
			if m[s[i:j+1]] {
				dp[j+1] = true
			}
		}
		if dp[n] == true {
			return true
		}
	}
	fmt.Println(dp)
	return dp[n] == true
}

//func minimumTotal(triangle [][]int) int {
//	m := len(triangle)
//	n := len(triangle[0])
//	dp := make([]int, n)
//}
