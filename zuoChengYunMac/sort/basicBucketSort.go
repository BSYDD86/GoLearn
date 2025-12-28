package main

//	func main() {
//		nums := []int{13, 17, 100, 27, 25}
//		fmt.Println(nums)
//		ints := bucketSort(nums)
//		fmt.Println(ints)
//		sort.Ints(nums)
//		fmt.Println(nums)
//	}
func basicBucketSort(arr []int) []int {
	if len(arr) == 0 {
		return arr
	}
	base := 10
	needNextTurn := true
	for needNextTurn {
		needNextTurn = false
		bucket := make([][]int, 10)
		for _, num := range arr {
			if !needNextTurn {
				if num >= base {
					needNextTurn = true
				}
			}
			digit := (num / (base / 10)) % 10
			bucket[digit] = append(bucket[digit], num)
		}
		arr = []int{}
		for _, b := range bucket {
			arr = append(arr, b...)
		}
		base *= 10
	}
	return arr
}
