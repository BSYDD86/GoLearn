package main

func main() {
	nums := []int{5, 2, 3, 1}
	sortArray(nums)
}

func sortArray(nums []int) []int {
	mergeSort(nums, 0, len(nums)-1)
	return nums
}

func mergeSort(nums []int, l int, r int) {
	if l == r {
		return
	}
	mid := (r-l)>>1 + l
	mergeSort(nums, l, mid)
	mergeSort(nums, mid+1, r)
	//merge
	target := nums[mid]
	temp := make([]int, r-l+1)
	tempL := 0
	tempR := r - l
	for i := l; i <= r; i++ {
		if nums[i] <= target {
			temp[tempL] = nums[i]
			tempL++
		} else {
			temp[tempR] = nums[i]
			tempR--
		}
	}
	for i, num := range temp {
		nums[i+l] = num
	}
}
