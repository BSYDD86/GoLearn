package main

import (
	"math"
)

func main() {
	res := [][]int{}
	res = append(res, []int{1, 2})

}

var (
	m map[byte]bool
)

func vowelStrings(words []string, left int, right int) int {
	res := 0
	m = map[byte]bool{'a': true, 'e': true, 'i': true, 'o': true, 'u': true}
	for _, word := range words {
		if checkLegal(word, left, right) {
			res++
		}
	}
	return res
}
func checkLegal(str string, l, r int) bool {
	if !m[str[l]] || !m[str[min(r, len(str)-1)]] {
		return false
	}
	return true
}

//func countGoodTriplets(arr []int, a int, b int, c int) int {
//res:=make([][]int,0)
//for i:=0;i<len(arr)-2;i++{
//	for j:=i+1;j<len(arr)-1;j++{
//		for k:=j+1;j<len(arr);j++{
//			if -a<=arr[i]-arr[j]&&arr[i]-arr[j]<=a{
//				if -b<=arr[j]-arr[k]&&arr[j]-arr[k]<=b{
//					if -c<=arr[i]-arr[k]&&arr[i]-arr[k]<=c{
//						res=append(res,[]int{arr[[i],arr[j],arr[k]})
//					}
//				}
//			}
//		}
//	}
//}
//return len(res)
//}

func maxScore(s string) int {
	oneCnt := 0
	for _, c := range s {
		if c == '1' {
			oneCnt++
		}
	}
	res := math.MinInt
	zeroCnt := 0
	for _, c := range s {
		if c == '1' {
			if oneCnt > 0 {
				oneCnt--
			}

		}
		if c == '0' {
			zeroCnt++
		}
		res = max(res, oneCnt+zeroCnt)
	}
	return res
}

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}
	res := 1
	for i := 0; i < n-1; i++ {
		for j := 1; j < n; j++ {
			if nums[i] < nums[j] {
				dp[j] = max(dp[j], dp[i]+1)
				res = max(res, dp[j])
			}
		}
	}
	return res
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt
	}
	dp[0] = 0
	for i := 0; i < amount; i++ {
		for _, j := range coins {
			if i+j <= amount {
				dp[i+j] = min(dp[i+j], dp[i]+1)
			}
		}
	}
	return dp[amount]
}

var (
	res []string
)

func dfs(left, right int, path []byte) {
	if left == 0 && right == 0 {
		res = append(res, string(path))
		return
	}
	if left >= right {
		dfs(left-1, right, append(path, '('))
		if right > 0 {
			dfs(left, right-1, append(path, ')'))
		}
	} else {
		if left > 0 {
			dfs(left-1, right, append(path, '('))
		}
		dfs(left, right-1, append(path, ')'))
	}
}
