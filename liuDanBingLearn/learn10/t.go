package main

func removeElement(nums []int, val int) int {
	var idx = 0
	var len = len(nums)

	for i := 0; i < len; i++ {
		if nums[i] != val {
			nums[idx] = nums[i]
			idx++
		}
	}
	return idx
}
