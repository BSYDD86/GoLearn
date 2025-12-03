package main

func main() {
	arr := []int{1, 3, 4, 2, 5}
	println(process(arr, 0, len(arr)-1))

}

// []
func process(arr []int, l int, r int) int {
	if r == l {
		return 0
	}
	mid := (r-l)/2 + l
	return process(arr, l, mid) +
		process(arr, mid+1, r) +
		merge(arr, l, mid, r)
}
func merge(arr []int, l int, mid int, r int) int {
	res := 0
	help := make([]int, r-l+1)
	idx := 0
	p1 := l
	p2 := mid + 1
	for p1 <= mid && p2 <= r {
		if arr[p1] < arr[p2] {
			res += arr[p1] * (r - p2 + 1)
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
		arr[l+i] = help[i]
	}
	return res
}
