package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	arr := []int{2, 4, 6, 5, 7}
	fmt.Println(minSwaps(arr))
}

func minSwaps(nums []int) int {
	odd := 0
	even := 0
	n := len(nums)
	for _, num := range nums {
		if num%2 == 1 {
			odd++
		} else {
			even++
		}
	}
	if n%2 == 0 {
		if odd != even {
			return -1
		}
	} else {
		temp := odd - even
		if temp < 0 {
			temp = -temp
		}
		if temp != 1 {
			return -1
		}
	}
	res := 0
	if n%2 == 1 {
		isOdd := true
		if even-odd == 1 {
			isOdd = false
		}
		for i := range nums {
			if isOdd {
				fmt.Println(i, ",", res)
				if i%2 == 1 {
					if nums[i]%2 == 1 {
						res += searchAndSwap(nums, i, isOdd)
					}
				} else {
					if nums[i]%2 == 0 {
						res += searchAndSwap(nums, i, isOdd)
					}
				}
			} else {
				fmt.Println(i, ",", res)
				if i%2 == 1 {
					if nums[i]%2 == 0 {
						res += searchAndSwap(nums, i, isOdd)
					}
				} else {
					if nums[i]%2 == 1 {
						res += searchAndSwap(nums, i, isOdd)
					}
				}
			}
		}
		return res
	} else {
		temp := make([]int, n)
		tempRes := 0
		copy(temp, nums)
		//odd,even
		for i := range nums {
			if i%2 == 1 {
				if nums[i]%2 != 1 {
					res += searchAndSwap(nums, i, true)
				}
			} else {
				if nums[i]%2 != 0 {
					res += searchAndSwap(nums, i, false)
				}
			}
		}
		//even,odd
		for i := range nums {
			if i%2 == 1 {
				if nums[i]%2 != 0 {
					tempRes += searchAndSwap(temp, i, false)
				}
			} else {
				if nums[i]%2 != 1 {
					tempRes += searchAndSwap(temp, i, true)
				}
			}
		}
		return min(res, tempRes)
	}
	return res
}

// swap baseIdx targetIdx
func searchAndSwap(nums []int, baseIdx int, isOdd bool) int {
	j := baseIdx + 1
	for j < len(nums) {
		if isOdd {
			if nums[j]%2 == 1 {
				break
			}
		} else {
			if nums[j]%2 == 0 {
				break
			}
		}
		j++
	}
	for i := j; i-1 != baseIdx; i-- {
		nums[i], nums[i-1] = nums[i-1], nums[i]
	}
	return j - baseIdx
}

func maxArea(coords [][]int) int64 {
	mX := map[int][][]int{}
	mY := map[int][][]int{}
	mXidx := map[int]bool{}
	mYidx := map[int]bool{}
	for _, coord := range coords {
		x := coord[0]
		y := coord[1]
		mX[x] = append(mX[x], []int{x, y})
		mY[y] = append(mY[y], []int{x, y})
		mXidx[x] = true
		mYidx[y] = true
	}
	res := -1
	xMin := math.MaxInt
	xMax := math.MinInt
	yMin := math.MaxInt
	yMax := math.MinInt
	for v := range mXidx {
		xMin = min(xMin, v)
		xMax = max(xMax, v)
	}
	for v := range mYidx {
		yMin = min(yMin, v)
		yMax = max(yMax, v)
	}
	if xMin == math.MaxInt ||
		xMax == math.MinInt ||
		yMin == math.MaxInt ||
		yMax == math.MinInt {
		return -1
	}
	for k := range mX {
		if len(mX[k]) == 1 {
			continue
		}
		iMin := math.MaxInt
		iMax := math.MinInt
		for _, v := range mX[k] {
			y := v[1]
			iMin = min(y, iMin)
			iMax = max(y, iMax)
		}
		temp := max(k-xMin, xMax-k)
		res = max(res, temp*(iMax-iMin))
	}

	for k := range mY {
		if len(mY[k]) == 1 {
			continue
		}
		iMin := math.MaxInt
		iMax := math.MinInt
		for _, v := range mY[k] {
			x := v[0]
			iMin = min(x, iMin)
			iMax = max(x, iMax)
		}
		temp := max(k-yMin, yMax-k)
		res = max(res, temp*(iMax-iMin))
	}
	return int64(res)
}

func minimumDeletions(word string, k int) int {
	n := len(word)
	arr := make([]int, 26)
	for _, c := range word {
		arr[c-'a']++
	}
	sort.Ints(arr)
	if k == 0 {
		return n - arr[25]
	}
	cnt := 0
	for i := 0; i < min(k, 26); i++ {
		cnt += arr[i]
	}
	res := cnt
	for i := k; i < 26; i++ {
		cnt += arr[i]
		cnt -= arr[i-k]
		res = max(res, cnt)
		fmt.Println("i is:", i, "cnt is:", cnt)
	}
	return n - res
}
