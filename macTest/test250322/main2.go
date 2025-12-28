package main

import "fmt"

func main() {
	fmt.Println("yes")
	nums := []int{2, 3, 4, 5, 2, 1, -1, 0}
	heapSort(nums)
	fmt.Println(nums)
	n := len(nums)
	res := [][]int{}
	st := []int{0}
	for len(st) > 0 {
		temp := []int{}
		tmepLen := len(st)
		for i := 0; i < tmepLen; i++ {
			temp = append(temp, nums[st[i]])
			if st[i]*2+1 < n {
				st = append(st, st[i]*2+1)
			}
			if st[i]*2+2 < n {
				st = append(st, st[i]*2+2)
			}
		}
		st = st[tmepLen:]
		res = append(res, temp)
	}
	fmt.Println(res)
}

func heapSort(arr []int) {
	for i := len(arr)/2 - 1; i >= 0; i-- {
		heapify(arr, len(arr), i)
	}
	//for i := len(arr) - 1; i >= 0; i-- {
	//	arr[i], arr[0] = arr[0], arr[i]
	//	heapify(arr, i, 0)
	//}
}
func heapify(arr []int, n int, i int) {
	l := 2*i + 1
	r := 2*i + 2
	tempIdx := i
	if l < n && arr[tempIdx] < arr[l] {
		tempIdx = l
	}
	if r < n && arr[tempIdx] < arr[r] {
		tempIdx = r
	}
	if i != tempIdx {
		arr[i], arr[tempIdx] = arr[tempIdx], arr[i]
		heapify(arr, n, tempIdx)
	}
}
