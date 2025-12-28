package main

func main() {

}

func countGood(nums []int, k int) int64 {
	cntArr := make([]int, len(nums))
	tempArr := make([]int, len(nums))
	l := 0
	r := 0
	for i, num := range nums {
		for l <= r {
			mid := (r-l)>>1 + l
			if cntArr[l] < num {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		cntArr[l] = num
		if r == l {
			r++
		}
		tempArr[i] = l
	}
	rArr := make([]int, len(nums))
	m := map[int]int{}
	for i := len(nums) - 1; i >= 0; i-- {
		rArr[i] = m[nums[i]]
		m[nums[i]]++
	}
	res := 0
	for i := range nums {
		if tempArr[i]*rArr[i] >= k {
			res++
		}
	}
	return int64(res)
}
