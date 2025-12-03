package main

import (
	"container/heap"
	"sort"
	"strings"
)

func sortVowels(s string) string {
	arr := []rune(s)
	mHeap := &myHeap{}
	heap.Init(mHeap)
	for _, v := range arr {
		if isVowel(v) {
			heap.Push(mHeap, v)
		}
	}
	for i, v := range arr {
		if isVowel(v) {
			arr[i] = heap.Pop(mHeap).(rune)
		}
	}
	sort.Slice(arr)

	return string(arr)

}

type myHeap []rune

func (m myHeap) Len() int {
	return len(m)
}

func (m myHeap) Less(i, j int) bool {
	return m[i] > m[j]
}

func (m *myHeap) Swap(i, j int) {
	(*m)[i], (*m)[j] = (*m)[j], (*m)[i]
}

func (m *myHeap) Push(x any) {
	(*m) = append((*m), x.(rune))
}

func (m *myHeap) Pop() any {
	res := (*m)[len(*m)-1]
	(*m) = (*m)[:len(*m)-1]
	return res
}

func isVowel(ch rune) bool {
	return strings.ContainsRune("aeiouAEIOU", ch)
}
