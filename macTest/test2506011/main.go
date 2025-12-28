package main

import (
	"container/heap"
	"fmt"
)

// IntHeap 是一个由整数组成的大顶堆
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

func main() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("最大值：%d\n", (*h)[0]) // 输出最大值

	// 依次弹出元素，验证顺序
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
