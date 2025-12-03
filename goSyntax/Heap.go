package main

import (
	"container/heap"
	"fmt"
)

func main() {
	mHeap := &myHeap{}
	heap.Init(mHeap)
	arr := []int{1, 3, 2}
	for _, entry := range arr {
		heap.Push(mHeap, entry)
	}
	for mHeap.Len() > 0 {
		fmt.Println(heap.Pop(mHeap))
	}
}

type myHeap []int

func (m myHeap) Len() int {
	return len(m)
}

// max heap
func (m myHeap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m *myHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}
func (m *myHeap) Push(x any) {
	(*m) = append((*m), x.(int))
}
func (m *myHeap) Pop() any {
	res := (*m)[len(*m)-1]
	(*m) = (*m)[:len(*m)-1]
	return res
}
