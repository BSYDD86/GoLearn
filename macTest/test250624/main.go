package main

import "fmt"

const mx = 3

var (
	res  [][]int
	path []int
)

func main() {
	path = []int{}
	res = [][]int{}
	dfs(1)
	fmt.Println(res)
}
func dfs(idx int) {
	if idx > mx {
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)
		return
	}
	path = append(path, idx)
	dfs(idx + 1)
	path = path[:len(path)-1]

}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
