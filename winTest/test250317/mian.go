package main

import "fmt"

func main() {
	//arr := []int{1, 2, 3, 4, 5}
	//test := arr[1:2]
	//fmt.Println(cap(test))
	//
	//arrByte := []byte{'('}
	//println(string(arrByte))
	//
	//println(generateParenthesis(3))
	i := partition("a")
	for _, j := range i {
		fmt.Println(j)
	}
}

var (
	dp   [][]bool
	path []string
	res  [][]string
)

func partition(s string) [][]string {
	n := len(s)
	dp = make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	for l := n - 1; l >= 0; l-- {
		for r := l; r < n; r++ {
			if s[l] == s[r] {
				if r-l <= 1 {
					dp[l][r] = true
				} else if dp[l+1][r-1] {
					dp[l][r] = true
				}
			}
		}
	}
	dfs(s, 0)
	return res
}
func dfs(s string, idx int) {
	if idx == len(s) {
		t := make([]string, len(path))
		copy(t, path)
		res = append(res, t)
		return
	}
	for i := idx; i < len(s); i++ {
		if dp[idx][i] {
			path = append(path, s[idx:i+1])
			dfs(s, i+1)
			path = path[:len(path)-1]
		}
	}
}

//var (
//	path []byte
//	res  []string
//)
//
//func generateParenthesis(n int) []string {
//	path = []byte{}
//	res = []string{}
//	dfs(n, n)
//	return res
//}
//
//func dfs(l int, r int) {
//	if l == 0 && r == 0 {
//		t := make([]byte, len(path))
//		copy(t, path)
//		res = append(res, string(t))
//		return
//	}
//	if l == r {
//		path = append(path, '(')
//		dfs(l-1, r)
//	} else if l <= r {
//		if l > 0 {
//			path = append(path, '(')
//			dfs(l-1, r)
//			path = path[:len(path)-1]
//		}
//		path = append(path, ')')
//		dfs(l, r-1)
//	}
//}
