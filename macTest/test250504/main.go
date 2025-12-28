package main

import "fmt"

func main() {
	i := -3
	i >>= 31
	i &= 1
	fmt.Println(i)

}

func halveArray(myNums []int) int {
	nums := make([]float64, len(myNums))
	for i := 0; i < len(nums); i++ {
		nums[i] = float64(myNums[i])
	}
	n := len(nums)
	temp := float64(0)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(nums, i, len(nums))
	}
	for i := range nums {
		temp += nums[i]
	}
	temp /= 2
	res := 0
	for temp > 0 {
		nums[0] /= 2
		temp -= nums[0]
		res++
		heapify(nums, 0, len(nums))
	}
	return res
}

func heapify(nums []float64, idx int, heapSize int) {
	best := idx
	l := 2*idx + 1
	r := 2*idx + 2
	if l < heapSize && nums[best] < nums[l] {
		best = l
	}
	if r < heapSize && nums[best] < nums[r] {
		best = r
	}
	if best != idx {
		nums[idx], nums[best] = nums[best], nums[idx]
		heapify(nums, best, heapSize)
	}
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
func (h *IntHeap) Pop() int {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
func (h *IntHeap) Top() int {
	return (*h)[0]
}
