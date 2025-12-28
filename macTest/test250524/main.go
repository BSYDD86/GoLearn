package main

import (
	"fmt"
)

func main() {
	word := "abcdeafdef"
	cnt := maxSubstrings(word)
	fmt.Println(cnt)
}

func maxSubstrings(word string) int {
	res := 0
	n := len(word)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		for j := i + 3; i <= n; i++ {
			if word[i-1] == word[j-1] {
				dp[j] = max(dp[j], dp[i-1]+1)
				res = max(res, dp[j])
			}
		}
	}
	return res
}

func findWordsContaining(words []string, x byte) []int {
	res := []int{}
cnt:
	for i, word := range words {
		for _, c := range word {
			if byte(c) == x {
				res = append(res, i)
				continue cnt
			}
		}
	}
	return res
}
