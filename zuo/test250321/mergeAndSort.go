package main

import "fmt"

func main() {
	nums := []int{2, -3, 1, 4, 5, 0}
	sort := mergeSort(nums)
	fmt.Println(sort)
}

// []
func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	l := mergeSort(arr[:mid])
	r := mergeSort(arr[mid:])
	return merge(l, r)
}
func merge(arr1 []int, arr2 []int) []int {
	res := make([]int, len(arr1)+len(arr2))
	i := 0
	j := 0

	idx := 0
	m := len(arr1)
	n := len(arr2)
	for i < m && j < n {
		if arr1[i] < arr2[j] {
			res[idx] = arr1[i]
			i++
		} else {
			res[idx] = arr2[j]
			j++
		}
		idx++
	}
	for i < m {
		res[idx] = arr1[i]
		i++
		idx++
	}
	for j < n {
		res[idx] = arr2[j]
		j++
		idx++
	}
	return res
}
