package main

func main() {
	//arr := []int{1, 3}
	//println(getIdxFromArr(arr, 2))
	arr := []int{1, 2, 3, 1}
	println(findPeakElement(arr))
}

func getIdxFromArr(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}
	l := 0
	r := len(arr) - 1
	for l <= r {
		mid := (r-l)/2 + l
		if arr[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}

func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	l := 0
	r := len(nums) - 1

	if nums[l] > nums[l+1] {
		return l
	}
	if nums[r] > nums[r-1] {
		return r
	}
	l++
	r--
	for l < r {
		mid := (r-l)/2 + l
		if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
		}
		if nums[mid] > nums[mid-1] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l
}
