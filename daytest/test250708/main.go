package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var (
	sum int
	res int
)

const mx = 1_000_000_009

func maxProduct(root *TreeNode) int {
	sum = dfsSum(root)
	res = 0
	dfs(root)
	return res
}
func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	temp := root.Val + dfs(root.Left) + dfs(root.Right)
	res = max(res, temp*(sum-temp)%mx)
	return temp
}
func dfsSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return dfsSum(root.Left) + dfsSum(root.Right) + root.Val
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
