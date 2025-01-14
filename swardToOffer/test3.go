package main

func main() {
	arr := []int{1, 1, 1}
	println(getMulty(arr))

}

func getMulty(arr []int) int {
	l := 1
	r := len(arr) - 1
	//[]
	for l <= r {
		mid := (r-l)/2 + l
		cnt := count(arr, l, mid)
		if l == r {
			if cnt > 1 {
				return l
			} else {
				break
			}

		}
		if cnt > r-l+1 {
			r = mid
		} else {
			l = mid + 1
		}

	}
	return -1

}
func count(arr []int, lIdx int, rIdx int) int {
	cnt := 0
	for _, temp := range arr {

		if temp >= lIdx && temp <= rIdx {
			cnt++
		}
	}
	return cnt

}
