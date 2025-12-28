package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeInfo struct {
	Nodes  int
	Height int
}

func isCompleteTree(root *TreeNode) bool {
	curTreeInfo := dfs(root)
	return curTreeInfo.Nodes == 1<<curTreeInfo.Height-1
}

func dfs(root *TreeNode) TreeInfo {
	if root == nil {
		return TreeInfo{0, 0}
	}
	l := dfs(root.Left)
	r := dfs(root.Right)
	return TreeInfo{l.Nodes + r.Nodes + 1, max(l.Height, r.Height)}
}
