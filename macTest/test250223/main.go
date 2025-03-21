package main

import "fmt"

func main() {
	arr := []int{3, 2, 1}
	nextPermutation(arr)
	fmt.Println(arr)
}

func nextPermutation(nums []int) {
	n := len(nums)
	for l := n - 2; l >= 0; l-- {
		for r := n - 1; r > l; r-- {
			if nums[r] > nums[l] {
				nums[l], nums[r] = nums[r], nums[l]
				reverse(nums, l+1, n-1)
				return
			}
		}
		reverse(nums, 0, n-1)
	}
}

// reverse
// []
func reverse(nums []int, l int, r int) {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	m := map[int][]int{}
	for _, p := range prerequisites {
		m[p[1]] = append(m[p[1]], p[0])
	}
	colors := make([]int, numCourses)
	//have chain
	var dfs func(int) bool
	dfs = func(i int) bool {
		colors[i] = 1
		for _, p := range m[i] {
			if colors[p] == 1 || colors[p] == 0 && dfs(p) {
				return true
			}

		}
		colors[i] = 2
		return false
	}
	for i := range m {
		if colors[i] == 0 && dfs(i) {
			return false
		}
	}
	return true
}

//func areOccurrencesEqual(s []byte) bool {
//	m := map[byte]int{}
//	for _, c := range s {
//		m[c]++
//	}
//	mSet := map[int]bool{}
//	for _, v := range m {
//		if mSet[v] {
//			return false
//		}
//		mSet[v] = true
//	}
//	return true
//}

//func maximalSquare( matrix [][]byte ) int {
//	m:=len(matrix)
//	n:=len(matrix[0])
//	dp:=make([][]int,m)
//	for i:=range dp{
//		dp[i]=make([]int,n)
//	}
//	res:=0
//	for i:=0;i<m;i++{
//		for j:=0;j<n;j++{
//			if matrix[i][j]=='1'{
//				if i==0||j==0{
//					dp[i][j]=1
//				}else{
//					dp[i][j]=max(dp[i-1][j],dp[i][j-1],dp[i-1][j-1])+1
//				}
//				res=max(res,dp[i][j])
//			}
//		}
//	}
//	return res*res
//}
