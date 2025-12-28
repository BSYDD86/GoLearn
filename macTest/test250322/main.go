package main

//func main() {
//	arr := []int{2, 3, 4, 5, 2, 1, -1, 0}
//	mergeSort(arr, 0, len(arr)-1)
//	fmt.Println(arr)
//}

// 归并
// []
func mergeSort(arr []int, l int, r int) {
	if r == l {
		return
	}
	mid := (r-l)/2 + l
	mergeSort(arr, l, mid)
	mergeSort(arr, mid+1, r)
	merge(arr, l, mid, r)
}
func merge(arr []int, l int, mid int, r int) {
	help := make([]int, r-l+1)
	p1 := l
	p2 := mid + 1
	idx := 0
	for p1 <= mid && p2 <= r {
		if arr[p1] < arr[p2] {
			help[idx] = arr[p1]
			p1++
		} else {
			help[idx] = arr[p2]
			p2++
		}
		idx++
	}
	for p1 <= mid {
		help[idx] = arr[p1]
		p1++
		idx++
	}
	for p2 <= r {
		help[idx] = arr[p2]
		p2++
		idx++
	}
	for i := 0; i < len(help); i++ {
		arr[i+l] = help[i]
	}
}
