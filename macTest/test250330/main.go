package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	res := ""
	str := "aab"
	res = res + str[0:1]
	fmt.Println(res)
	pow := int(math.Pow(10, 2))
	fmt.Println(pow)
	fmt.Println(reflect.TypeOf(pow))
}

var (
	cost []int64
	res  int64
)

func minOperations(nums []int, x int, k int) int64 {
	res = math.MaxInt64
	cost = make([]int64, len(nums))
	for i := 0; i-1+x < len(nums); i++ {
		cost[i] = minCost(nums, i, i-1+x)
	}
	dfs(cost, 0, x, k, 0)
	return res
}
func dfs(cost []int64, idx int, x int, k int, sum int64) {
	if k == 0 {
		res = min(res, sum)
		return
	}
	for i := idx; i-1+x*k < len(cost); i++ {
		dfs(cost, i+x, x, k-1, sum+cost[i])
	}
}
func minCost(arr []int, l int, r int) int64 {
	temp := 0
	for i := l; i <= r; i++ {
		temp += arr[i]
	}
	base := temp / (r - l + 1)
	baseCnt := 0
	baseSmall := base - 1
	baseSmallCnt := 0
	baseBig := base + 1
	baseBigCnt := 0
	for i := l; i <= r; i++ {
		temp = arr[i]
		if temp > base {
			baseCnt += (temp - base)
		} else {
			baseCnt += (base - temp)
		}

		if temp > baseSmall {
			baseSmallCnt += (temp - baseSmall)
		} else {
			baseSmallCnt += (baseSmall - temp)
		}

		if temp > baseBig {
			baseBigCnt += (temp - baseBig)
		} else {
			baseBigCnt += (baseBig - temp)
		}
	}
	return int64(min(baseCnt, baseSmallCnt, baseBigCnt))
}

func longestPalindrome(s string, t string) int {
	sArr := []byte(s)
	tArr := []byte(t)
	m := len(sArr)
	n := len(tArr)
	for i := 0; i < n/2; i++ {
		tArr[i], tArr[n-1-i] = tArr[n-1-i], tArr[i]
	}
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if sArr[i-1] == tArr[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	temp := dp[m][n]
	if temp == 0 {
		return 1
	}
	temp *= 2
	if temp < m && temp < n {
		temp += 1
	}
	return temp
}
