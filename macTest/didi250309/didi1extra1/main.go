package main

import "fmt"

var (
	res  int
	base int
)

func main() {

	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		res = 0
		base = 0
		var t int
		fmt.Scan(&t)
		arr := make([]int, t)
		for j := 0; j < t; j++ {
			fmt.Scan(&arr[j])
		}
		for j := 0; j < t; j++ {
			base += getCnt(arr, j)
		}

		dfs(arr, []int{}, 0)
		fmt.Println(res % 998244353)
	}
}
func dfs(arr []int, path []int, idx int) {
	if idx == len(arr) {
		if len(path) == 0 {
			return
		}

		temp := 0
		for i := 0; i < len(path); i++ {
			temp += getCnt(path, i)
		}

		if temp == base {
			res++
		}
		return
	}
	dfs(arr, path, idx+1)
	path = append(path, arr[idx])
	dfs(arr, path, idx+1)
	path = path[:len(path)-1]
}
func getCnt(arr []int, idx int) int {
	cnt := 0
	temp := arr[idx]
	for i := 0; i < idx; i++ {
		if arr[i] > temp {
			cnt++
		}
	}
	for i := idx + 1; i < len(arr); i++ {
		if arr[i] < temp {
			cnt++
		}
	}
	return cnt
}
