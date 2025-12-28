package main

import (
	"math"
	"sort"
)

func main() {
	
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	iMin, res := dfs(root)
	if iMin == math.MaxInt || res == math.MaxInt {
		return -1
	}
	return res
}

//[fisrtMin,secondMin]

func dfs(root *TreeNode) (int, int) {
	if root == nil {
		return math.MaxInt, math.MaxInt
	}
	lmin, lSecondMin := dfs(root.Left)
	rmin, rSecondMin := dfs(root.Right)
	temp := []int{lmin, rmin, lSecondMin, rSecondMin}
	sort.Ints(temp)
	return temp[0], temp[1]
}
