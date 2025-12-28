package main

import "fmt"

func main() {
	words := []string{"lr", "h"}
	groups := []int{0, 0}
	fmt.Println(getLongestSubsequence(words, groups))
}

var (
	res  []string
	path []string
	pre  int
)

func getLongestSubsequence(words []string, groups []int) []string {
	res = []string{}
	path = []string{}
	pre = -1
	dfs(words, groups, 0)
	return res
}
func dfs(words []string, groups []int, idx int) {
	if idx == len(words) {
		if len(path) > len(res) {
			temp := make([]string, len(path))
			copy(temp, path)
			res = temp
		}
		return
	}
	for i := idx; i < len(words); i++ {
		if groups[i] != pre {
			path = append(path, words[i])
			pre = groups[i]
			dfs(words, groups, i+1)
			if idx == 0 {
				pre = -1
			} else {
				pre ^= groups[i]
			}
			path = path[:len(path)-1]
		}
	}
}
