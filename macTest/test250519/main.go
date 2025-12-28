package main

import fmt "fmt"

var (
	res  [][]int
	path []int
)

func main() {
	res = [][]int{}
	path = []int{}
	n := 4
	k := 2
	dfs(1, k, n)
	fmt.Println(res)
}

func dfs(idx int, k int, end int) {
	if len(path) == k {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)
		return
	}
	for i := idx; i+k-len(path)-1 <= end; i++ {
		path = append(path, i)
		dfs(i+1, k, end)
		path = path[:len(path)-1]
	}
}
