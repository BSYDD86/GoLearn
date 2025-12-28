package main

import (
	"fmt"
)

func main() {
	//response := [][]string{{"gk", "otg", "ia", "otg", "qs", "cwtk"}, {"i", "otg", "otg", "otg", "otg", "otg", "otg"}}
	//println(findCommonResponse(response))
	//arr := [][]int{
	//	{0, 1, 2},
	//	{0, 2, 3},
	//	{1, 3, 4},
	//	{1, 4, 5},
	//	{2, 5, 2},
	//	{4, 6, 3},
	//	{5, 7, 4},
	//}
	//res := baseUnitConversions(arr)
	//fmt.Println(res)

	arr := []int{2, 3}
	temp := 0
	arr = append([]int{temp}, arr...)
}

func findCommonResponse(responses [][]string) string {
	m := map[string]int{}
	cnt := 0
	res := ""
	for _, response := range responses {
		mTemp := map[string]bool{}
		for _, str := range response {
			if !mTemp[str] {
				m[str]++
				if m[str] > cnt {
					cnt = m[str]
					res = str
				} else if cnt == m[str] {
					res = min(res, str)
				}
				mTemp[str] = true
			}

		}
	}
	fmt.Println(m)
	return res
}

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	pre := -1
	res := 0
	iMax := -1
	for i, num := range nums {
		if num > right {
			pre = i
		} else {
			iMax = i
		}
		if iMax-pre > 0 {
			myPrint(nums, pre+1, iMax)
		}
		res += max((iMax - pre), 0)
	}
	return res
}

// []
func myPrint(nums []int, l int, r int) {
	temp := make([]int, r-l)
	for i := l; i <= r; i++ {
		temp = append(temp, nums[i])
	}
	fmt.Println(temp)
	return
}

type test struct {
	age    int
	gender string
}

func baseUnitConversions(conversions [][]int) []int {
	g := map[int]int{}
	for _, conv := range conversions {
		target := conv[1]
		source := conv[0]
		g[target] = source
	}
	n := len(conversions)
	res := make([]int, n+1)
	res[0] = 1
	var dfs func(idx int) int
	dfs = func(idx int) int {
		if res[idx] != 0 {
			return res[idx]
		} else {
			source := g[idx]
			factor := conversions[source][2]
			return dfs(g[idx]) * factor
		}
	}
	for _, conv := range conversions {
		source := conv[0]
		target := conv[1]
		factor := conv[2]
		res[target] = factor * dfs(source) % (1000000000 + 7)
	}
	return res
}

type MapHeap struct {
	heap []int
}

func (h *MapHeap) Push(val int) {
	h.heap = append(h.heap, val)

}
func (h *MapHeap) HeapInsert(idx int) {
	parentIdx := idx/2 - 1
	if parentIdx < 0 || h.heap[parentIdx] > h.heap[idx] {
		return
	}
	h.heap[parentIdx], h.heap[idx] = h.heap[idx], h.heap[parentIdx]
}

//10
//-2
//-1
