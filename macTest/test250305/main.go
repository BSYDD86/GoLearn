package main

//
//import (
//	"fmt"
//	"math"
//)
//
//func main() {
//	arr1 := []int{1, 3}
//	arr2 := []int{2}
//	fmt.Println(findMedianSortedArrays(arr1, arr2))
//}
//
//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//	//（）
//	if len(nums1) > len(nums2) {
//		nums1, nums2 = nums2, nums1
//	}
//	m := len(nums1)
//	n := len(nums2)
//	l := -1
//	r := m
//	for l < r {
//		i := (r-l)/2 + l
//		j := (m+n+1)/2 - i
//		if j >= n || nums1[i] <= nums2[j+1] {
//			l = i
//		} else {
//			r = j
//		}
//	}
//	i := (r-l)/2 + l
//	j := (m+n+1)/2 - i
//	minNums1 := math.MinInt
//	maxNums1 := math.MaxInt
//	minNums2 := math.MinInt
//	maxNums2 := math.MaxInt
//	if i >= 0 {
//		minNums1 = nums1[i]
//	}
//	if i+1 < m {
//		maxNums1 = nums1[i+1]
//	}
//	if j >= 0 {
//		minNums2 = nums2[j]
//	}
//	if j+1 < n {
//		maxNums2 = nums2[j]
//	}
//	return float64(max(minNums1, minNums2)+min(maxNums1, maxNums2)) / 2.0
//}
//
////// 回溯函数，用于生成所有可能的组合
////func backtrack(arr []int, start int, path []int, result *[][]int) {
////	// 把当前组合加入到结果集中
////	combination := make([]int, len(path))
////	copy(combination, path)
////	*result = append(*result, combination)
////
////	// 回溯遍历所有选择
////	for i := start; i < len(arr); i++ {
////		// 做选择
////		path = append(path, arr[i])
////		// 递归
////		backtrack(arr, i+1, path, result)
////		// 撤销选择
////		path = path[:len(path)-1]
////	}
////}
////
////// 主函数，生成组合
////func combine(arr []int) [][]int {
////	var result [][]int
////	backtrack(arr, 0, []int{}, &result)
////	return result
////}
////
////func main() {
////	arr := []int{1, 2, 3}
////	combinations := combine(arr)
////	for _, combination := range combinations {
////		fmt.Println(combination)
////	}
////}
//
////var (
////	temp int
////	res  int
////)
//
////func main() {
////	var m int
////	var n int
////	fmt.Scan(&m, &n)
////	arr := make([][]byte, m)
////	var strTemp string
////	for i := range arr {
////		fmt.Scan(&strTemp)
////		arr[i] = make([]byte, n)
////		for j := 0; j < n; j++ {
////			arr[i][j] = byte(strTemp[j])
////		}
////	}
////	dp := make([][]int, m)
////	for i := range dp {
////		dp[i] = make([]int, n)
////	}
////	for i := 0; i < m; i++ {
////		for j := 0; j < n; j++ {
////			dp[i][j] = math.MaxInt
////		}
////	}
////	dp[0][0] = 0
////	for i := 0; i < m; i++ {
////		for j := 0; j < n; j++ {
////			if dp[i][j] == math.MaxInt {
////				continue
////			}
////			c := getConvert(arr[i][j])
////			//down
////			if i+1 >= m || arr[i+1][j] == c {
////			} else {
////				dp[i+1][j] = min(dp[i][j]+1, dp[i+1][j])
////			}
////			//right
////			if j+1 >= n || arr[i][j+1] == c {
////			} else {
////				dp[i][j+1] = min(dp[i][j]+1, dp[i][j+1])
////			}
////		}
////	}
////	if dp[m-1][n-1] == math.MaxInt {
////		fmt.Println(-1)
////	} else {
////		fmt.Println(dp[m-1][n-1])
////	}
////}
////
////func getConvert(c byte) byte {
////	if c == 'r' {
////		return 'd'
////	}
////	if c == 'e' {
////		return 'r'
////	}
////	if c == 'd' {
////		return 'e'
////	}
////	return 'a'
////}
////
////func min(a, b int) int {
////	if a > b {
////		return b
////	} else {
////		return a
////	}
////}
////
////var (
////	temp int
////	res  int
////)
//
////func main() {
////	n := 0
////	k := 0
////	fmt.Scan(&n, &k)
////	arr := make([]int, n)
////	for i := range arr {
////		fmt.Scan(&arr[i])
////	}
////	for i := 0; i < len(arr); i++ {
////		temp = arr[i]
////		dfs(arr, i+1, k-1)
////	}
////	fmt.Println(res)
////}
////func dfs(arr []int, idx int, cnt int) {
////	if idx == len(arr) && cnt == 0 {
////		res = max(res, temp)
////		return
////	}
////	if idx >= len(arr) || cnt <= 0 {
////		return
////	}
////	for i := idx; i < len(arr); i++ {
////		dfs(arr, i+1, cnt)
////		temp = temp & arr[i]
////		dfs(arr, i+1, cnt-1)
////		temp = temp & arr[i]
////	}
////}
////
////func max(a, b int) int {
////	if a > b {
////		return a
////	} else {
////		return b
////	}
////}
//
////
////func main() {
////	var n int
////	var str string
////	var res string
////	fmt.Scan(&n, &str)
////	pre := 0
////	m := map[string]int{}
////	for i := 0; i+n <= len(str); i++ {
////		m[str[i:i+n]]++
////	}
////	for k, v := range m {
////		if v > pre {
////			res = k
////			pre = v
////		} else if v == pre {
////			if res > k {
////				res = k
////			}
////		}
////	}
////	fmt.Println(res)
////}
//
////func main() {
////	a := 0
////	b := 0
////	fmt.Scan(&a, &b)
////	fmt.Println(a)
////	fmt.Println(b)
////	res := 0
////	for i := a; i >= (a - b + 1); i-- {
////		res += i
////	}
////	fmt.Println(res)
////}
