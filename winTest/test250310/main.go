package main

import "slices"

//	func main() {
//		arr:=]
//
// }
func moveZeroes(nums []int) {
	n := len(nums)
	idx := 0
	for _, num := range nums {
		if num != 0 {
			nums[idx] = num
			idx++
		}
	}
	for i := idx; i < n; i++ {
		nums[idx] = 0
	}
}

func groupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	for _, str := range strs {
		arr := []byte(str)
		slices.Sort(arr)
		temp := string(arr)
		m[temp] = append(m[temp], str)
	}
	res := [][]string{}
	for _, v := range m {
		res = append(res, v)
	}
	return res
}
