package main

//var (
//	dp   [][]bool
//	path []string
//	res  [][]string
//)
//
//func partition(s string) [][]string {
//	m := len(s)
//	dp = make([][]bool, m)
//	path = make([]string, 0)
//	res = make([][]string, 0)
//
//	for i, _ := range dp {
//		dp[i] = make([]bool, m)
//	}
//	for i := m - 1; i >= 0; i++ {
//		for j := i; j < m; j++ {
//			if s[i] == s[j] {
//				if j-i <= 1 {
//					dp[i][j] = true
//				} else if dp[i+1][j-1] {
//					dp[i][j] = true
//				}
//			}
//		}
//	}
//	dfs(s, 0)
//	return res
//}
//
//// []
//func dfs(s string, idx int) {
//	if idx == len(s) {
//		t := make([]string, len(path))
//		copy(t, path)
//		res = append(res, t)
//		return
//	}
//	for i := idx; i < len(s); i++ {
//		if dp[idx][i] {
//			path = append(path, s[idx:i+1])
//			dfs(s, i+1)
//			path = path[:len(path)-1]
//		}
//	}
//}
//
//func main() {
//	s := "aab"
//	println(partition(s))
//}
