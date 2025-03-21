package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

var (
	path []byte
	res  []string
)

func generateParenthesis(n int) []string {
	path = []byte{}
	res = []string{}
	dfs(n, n)
	return res
}
func dfs(lCnt int, rCnt int) {
	if lCnt == 0 && rCnt == 0 {
		t := make([]byte, len(path))
		copy(t, path)
		res = append(res, string(t))
		return
	}
	if lCnt == rCnt {
		path = append(path, '(')
		dfs(lCnt-1, rCnt)
		path = path[:len(path)-1]
	}
	if rCnt > lCnt {
		path = append(path, '(')
		dfs(lCnt-1, rCnt)
		path = path[:len(path)-1]
		path = append(path, ')')
		dfs(lCnt, rCnt-1)
		path = path[:len(path)-1]
	}
}

func max(i ...int) int {
	t := i[0]
	for n := range i {
		if n < t {
			t = n
		}
	}
	return t
}
