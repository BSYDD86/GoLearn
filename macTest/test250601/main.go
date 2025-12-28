package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {
	dp := [][]int{{1, 8}, {3, -2}}
	fmt.Println(minAbsDiff(dp, 2))
}
func minAbsDiff(grid [][]int, k int) [][]int {
	res := math.MaxInt
	rows := len(grid)
	cols := len(grid[0])
	dp := make([][]int, rows-k+1)
	for i := range dp {
		dp[i] = make([]int, cols-k+1)
	}
	for i := 0; i < rows-k+1; i++ {
		for j := 0; j < cols-k+1; j++ {
			res = math.MaxInt
			h := &IntHeap{}
			heap.Init(h)
			for row := i; row < i+k; row++ {
				for col := j; col < j+k; col++ {
					heap.Push(h, grid[row][col])
				}
			}
			fmt.Println(h)
			pre := h.Pop()
			for h.Len() > 0 {
				temp := h.Pop()
				cur := pre.(int) - temp.(int)
				if cur < 0 {
					cur *= (-1)
				}
				pre = temp
				res = min(res, cur)
			}
			dp[i][j] = res
		}
	}
	return dp
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] } // 大顶堆比较
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push 往堆里添加元素
func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop 从堆中移除并返回堆顶元素
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
