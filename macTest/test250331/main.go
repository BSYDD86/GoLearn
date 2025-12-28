package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	arr := []int{13, 17, 100, 27, 25}
	mergeSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func morrisTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	var (
		morrisRight *TreeNode
	)
	cur := root
	for cur != nil {
		morrisRight = cur.Left
		if morrisRight != nil {
			for morrisRight.Right != nil && morrisRight.Right != cur {
				morrisRight = morrisRight.Right
			}
			if morrisRight.Right == nil {
				morrisRight.Right = cur
				cur = cur.Left
				continue
			} else {
				morrisRight.Right = nil
			}
		}
		cur = cur.Right
	}
}

// []
func mergeSort(arr []int, l int, r int) {
	if l >= r {
		return
	}
	mid := (r-l)/2 + l
	mergeSort(arr, l, mid)
	mergeSort(arr, mid+1, r)
	help := make([]int, r-l+1)
	cur := 0
	i := l
	j := mid + 1
	for i <= mid && j <= r {
		if arr[i] <= arr[j] {
			help[cur] = arr[i]
			i++
		} else {
			help[cur] = arr[j]
			j++
		}
		cur++
	}
	for i <= mid {
		help[cur] = arr[i]
		i++
		cur++
	}
	for j <= r {
		help[cur] = arr[j]
		j++
		cur++
	}
	for k := 0; k < len(help); k++ {
		arr[k+l] = help[k]
	}
}

func quickSort(arr []int, l int, r int) {
	if l >= r {
		return
	}
	idx := partition(arr, l, r)
	quickSort(arr, l, idx-1)
	quickSort(arr, idx+1, r)
}

// []
func partition(arr []int, l int, r int) int {
	randIdx := rand.Intn(r-l) + l
	arr[randIdx], arr[r] = arr[r], arr[randIdx]
	target := arr[r]
	preIdx := l
	for i := l; i < r; i++ {
		if arr[i] < target {
			arr[i], arr[preIdx] = arr[preIdx], arr[i]
			preIdx++
		}
	}
	arr[r], arr[preIdx] = arr[preIdx], arr[r]
	return preIdx
}

func heapSort(arr []int) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}
func heapify(arr []int, heapSize int, idx int) {
	if idx >= heapSize {
		return
	}
	l := 2*idx + 1
	r := 2*idx + 2
	tempIdx := idx
	if l < heapSize && arr[tempIdx] < arr[l] {
		tempIdx = l
	}
	if r < heapSize && arr[tempIdx] < arr[r] {
		tempIdx = r
	}
	if tempIdx != idx {
		arr[idx], arr[tempIdx] = arr[tempIdx], arr[idx]
		heapify(arr, heapSize, tempIdx)
	}
}
func bucketSort(arr []int, l int, r int, digit int) {
	bucket := make([]int, r-l+1)
	radix := 10
	for d := 1; d <= digit; d++ {
		count := make([]int, radix)
		for _, num := range arr {
			j := getDigit(num, d)
			count[j]++
		}
		for i := 1; i < len(count); i++ {
			count[i] = count[i-1] + count[i]
		}
		for i := r; i >= l; i-- {
			j := getDigit(arr[i], d)
			bucket[count[j]-1] = arr[i]
			count[j]--
		}
		arrIdx := l
		for _, num := range bucket {
			arr[arrIdx] = num
			arrIdx++
		}
	}
}
func getDigit(num int, digit int) int {
	pow := math.Pow(10, float64(digit-1))
	return (num / int(pow)) % 10
}
